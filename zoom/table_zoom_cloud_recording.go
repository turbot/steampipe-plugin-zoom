package zoom

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/himalayan-institute/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableZoomCloudRecording(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_cloud_recording",
		Description: "Meetings and webinars recorded to the cloud.",
		List: &plugin.ListConfig{
			Hydrate:    listCloudRecording,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		// TODO - SDK does not yet support Get
		Columns: []*plugin.Column{
			// Top columns
			{Name: "uuid", Type: proto.ColumnType_STRING, Description: "Unique Meeting Identifier. Each instance of the meeting will have its own UUID."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "Meeting ID - also known as the meeting number."},
			{Name: "topic", Type: proto.ColumnType_STRING, Description: "Recording topic."},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("StartTime").Transform(timeToTimestamp), Description: "Recording start time in GMT/UTC. Start time will not be returned if the recording is an instant recording."},
			// Other columns
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("HostID"), Description: "ID of the user who is set as the host of the recording."},
			{Name: "duration", Type: proto.ColumnType_INT, Transform: transform.FromField("Duration"), Description: "Recording duration."},
			{Name: "total_size", Type: proto.ColumnType_INT, Transform: transform.FromField("TotalSize"), Description: "Total size of the recording."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Recording Types: 1 - Instant recording. 2 - Scheduled recording. 3 - Recurring recording with no fixed time. 8 - Recurring recording with fixed time."},
			{Name: "share_url", Type: proto.ColumnType_STRING, Description: "Share URL for the recording."},
			{Name: "recording_count", Type: proto.ColumnType_INT, Transform: transform.FromField("RecordingCount"), Description: "Number of recording files returned in the response of this API call."},
			{Name: "recording_files", Type: proto.ColumnType_JSON, Description: "List of recording file."},
			{Name: "settings", Type: proto.ColumnType_JSON, Hydrate: getCloudRecordingSettings, Transform: transform.FromValue(), Description: "Settings for the recording."},
			// Common columns
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "Unique Identifier of the user account."},
		},
	}
}

func listCloudRecording(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_cloud_recording.listCloudRecording", "connection_error", err)
		return nil, err
	}
	keyQuals := d.KeyColumnQuals
	hostID := keyQuals["user_id"].GetStringValue()
	pageSize := 300
	opts := zoom.ListAllRecordingsOptions{
		UserID:   hostID,
		PageSize: &pageSize,
	}

	// Zoom only allows 30 days worth of recordings at a time. The default at the API is for the last
	// one day of recordings, which is not very helpful.
	// Approach:
	// * Respect the start_time for providing a range (watch out for caching!)
	// * By default, provide 30 days of recordings.
	quals := d.QueryContext.GetQuals()
	if quals["start_time"] != nil {
		for _, q := range quals["start_time"].Quals {
			ts := ptypes.TimestampString(q.Value.GetTimestampValue())
			switch q.GetStringValue() {
			case ">", ">=":
				opts.From = ts[0:10]
			case "<", "<=":
				opts.To = ts[0:10]
			}
		}
	} else {
		opts.From = time.Now().UTC().AddDate(0, -30, 0).Format("2006-01-02")
	}

	for {
		result, err := conn.ListAllRecordings(opts)
		if err != nil {
			if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
				// Host not found
				return nil, nil
			}
			plugin.Logger(ctx).Error("zoom_cloud_recording.listCloudRecording", "query_error", err)
			return nil, err
		}
		for _, i := range result.Meetings {
			d.StreamListItem(ctx, i)
		}
		if result.NextPageToken == "" {
			break
		}
		opts.NextPageToken = result.NextPageToken
	}
	return nil, nil
}

func getCloudRecordingSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_cloud_recording.getCloudRecordingSettings", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := int(quals["id"].GetInt64Value())
	if meeting, ok := h.Item.(zoom.CloudRecordingMeeting); ok {
		id = meeting.ID
	}
	opts := zoom.GetMeetingRecordingSettingsOptions{
		MeetingID: id,
	}
	result, err := conn.GetMeetingRecordingSettings(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 3001 {
			// Not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_cloud_recording.getCloudRecordingSettings", "query_error", err)
		return nil, err
	}
	return result, nil
}
