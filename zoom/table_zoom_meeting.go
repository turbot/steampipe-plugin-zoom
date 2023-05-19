package zoom

import (
	"context"

	"github.com/bigdatasourav/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableZoomMeeting(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_meeting",
		Description: "Scheduled meetings for the host. Instant meetings are not returned.",
		List: &plugin.ListConfig{
			Hydrate:    listMeeting,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Get: &plugin.GetConfig{
			Hydrate:        getMeeting,
			KeyColumns:     plugin.SingleColumn("id"),
			MaxConcurrency: 5,
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("HostID"), Description: "ID of the user who is set as the host of the meeting."},
			{Name: "id", Type: proto.ColumnType_INT, Description: "Meeting ID, also known as the meeting number."},
			{Name: "topic", Type: proto.ColumnType_STRING, Description: "Meeting topic."},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("StartTime").Transform(timeToTimestamp), Description: "Meeting start time in GMT/UTC. Start time will not be returned if the meeting is an instant meeting."},
			// Other columns
			{Name: "agenda", Type: proto.ColumnType_STRING, Hydrate: getMeeting, Transform: transform.FromField("Agenda"), Description: "Meeting description."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToTimestamp), Description: "Time when the meeting was created."},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "Meeting duration."},
			{Name: "encrypted_password", Type: proto.ColumnType_STRING, Hydrate: getMeeting, Description: "Encrypted passcode for third party endpoints (H323/SIP)."},
			{Name: "h323_password", Type: proto.ColumnType_STRING, Hydrate: getMeeting, Description: "H.323/SIP room system passcode."},
			{Name: "join_url", Type: proto.ColumnType_STRING, Description: "URL using which participants can join a meeting."},
			{Name: "occurrences", Type: proto.ColumnType_JSON, Hydrate: getMeeting, Description: "Array of occurrence objects."},
			{Name: "password", Type: proto.ColumnType_STRING, Hydrate: getMeeting, Description: "Meeting passcode."},
			{Name: "pmi", Type: proto.ColumnType_STRING, Hydrate: getMeeting, Transform: transform.FromField("PMI"), Description: "Personal meeting ID."},
			{Name: "recurrence", Type: proto.ColumnType_JSON, Hydrate: getMeeting, Description: "Recurrence details."},
			{Name: "settings", Type: proto.ColumnType_JSON, Hydrate: getMeeting, Description: "Meeting settings."},
			{Name: "start_url", Type: proto.ColumnType_STRING, Description: "The start_url of a meeting is a URL using which a host or an alternative host can start the Meeting."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Meeting status: waiting or started."},
			{Name: "timezone", Type: proto.ColumnType_STRING, Description: "Timezone to format the meeting start time."},
			{Name: "tracking_fields", Type: proto.ColumnType_JSON, Hydrate: getMeeting, Description: "Tracking fields."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Meeting Types: 1 - Instant meeting. 2 - Scheduled meeting. 3 - Recurring meeting with no fixed time. 8 - Recurring meeting with fixed time."},
			{Name: "uuid", Type: proto.ColumnType_STRING, Description: "Unique Meeting ID. Each meeting instance will generate its own Meeting UUID."},
		}),
	}
}

func listMeeting(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	userID := quals["user_id"].GetStringValue()
	pageSize := 300
	opts := zoom.ListMeetingsOptions{
		HostID: userID,
		// "scheduled" includes all valid past meetings (unexpired), live meetings
		// and upcoming scheduled meetings.
		Type:     "scheduled",
		PageSize: &pageSize,
	}
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil { // check if JWT creds is set
		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_meeting.connect.listMeeting", "connection_error", err)
			return nil, err
		}
		for {
			result, err := conn.ListMeetings(opts)
			if err != nil {
				if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
					// Host not found
					return nil, nil
				}
				plugin.Logger(ctx).Error("zoom_meeting.connect.listMeeting", "query_error", err)
				return nil, err
			}
			for _, i := range result.Meetings {
				d.StreamListItem(ctx, i)
			}
			if result.NextPageToken == "" {
				break
			}
			opts.NextPageToken = &result.NextPageToken
		}
	} else { // check if server-to-server oauth creds is set
		conn, err := connectOAuth(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_meeting.connectOAuth.listMeeting", "connection_error", err)
			return nil, err
		}
		for {
			result, err := conn.ListMeetings(opts)
			if err != nil {
				if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
					// Host not found
					return nil, nil
				}
				plugin.Logger(ctx).Error("zoom_meeting.connectOAuth.listMeeting", "query_error", err)
				return nil, err
			}
			for _, i := range result.Meetings {
				d.StreamListItem(ctx, i)
			}
			if result.NextPageToken == "" {
				break
			}
			opts.NextPageToken = &result.NextPageToken
		}
	}
	return nil, nil
}

func getMeeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	id := int(quals["id"].GetInt64Value())
	if meeting, ok := h.Item.(zoom.ListMeeting); ok {
		id = meeting.ID
	}
	opts := zoom.GetMeetingOptions{
		MeetingID: id,
	}
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil { // check if JWT creds is set
		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_meeting.connect.getMeeting", "connection_error", err)
			return nil, err
		}
		getMeetingWithRetry := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
			return conn.GetMeeting(opts)
		}
		result, err := plugin.RetryHydrate(ctx, d, h, getMeetingWithRetry, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			if e, ok := err.(*zoom.APIError); ok && e.Code == 3001 {
				// Meeting not found
				return nil, nil
			}
			plugin.Logger(ctx).Error("zoom_meeting.connect.getMeeting", "query_error", err)
			return nil, err
		}
		return result, nil
	} else { // check if server-to-server oauth creds is set
		conn, err := connectOAuth(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_meeting.connectOAuth.getMeeting", "connection_error", err)
			return nil, err
		}

		getMeetingWithRetry := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
			return conn.GetMeeting(opts)
		}
		result, err := plugin.RetryHydrate(ctx, d, h, getMeetingWithRetry, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})

		if err != nil {
			if e, ok := err.(*zoom.APIError); ok && e.Code == 3001 {
				// Meeting not found
				return nil, nil
			}
			plugin.Logger(ctx).Error("zoom_meeting.connectOAuth.getMeeting", "query_error", err)
			return nil, err
		}
		return result, nil
	}
}
