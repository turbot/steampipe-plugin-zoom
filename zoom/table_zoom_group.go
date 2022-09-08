package zoom

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableZoomGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_group",
		Description: "A Zoom account can have one or more groups.",
		List: &plugin.ListConfig{
			Hydrate: listGroup,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getGroup,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Group ID."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Group name."},
			{Name: "total_members", Type: proto.ColumnType_INT, Transform: transform.FromField("TotalMembers"), Description: "Total number of members in the group."},
		}),
	}
}

func listGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_group.listGroup", "connection_error", err)
		return nil, err
	}
	result, err := conn.ListGroups()
	if err != nil {
		plugin.Logger(ctx).Error("zoom_group.listGroup", "query_error", err)
		return nil, err
	}
	for _, i := range result.Groups {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_group.getGroup", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	opts := zoom.GetGroupOpts{
		ID: id,
	}
	result, err := conn.GetGroup(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// Group not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_group.getGroup", "query_error", err)
		return nil, err
	}
	return result, nil
}
