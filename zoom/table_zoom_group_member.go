package zoom

import (
	"context"

	"github.com/turbot/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableZoomGroupMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_group_member",
		Description: "Members of a Zoom group.",
		List: &plugin.ListConfig{
			Hydrate:    listGroupMember,
			KeyColumns: plugin.SingleColumn("group_id"),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "group_id", Type: proto.ColumnType_STRING, Hydrate: groupIDString, Transform: transform.FromValue(), Description: "Group ID."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID"), Description: "User ID of the member."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the member."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "First name of the member."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "Last name of the member."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the member."},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "Department of the member."},
		}),
	}
}

func listGroupMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_group.listGroupMember", "connection_error", err)
		return nil, err
	}
	groupID := d.EqualsQuals["group_id"].GetStringValue()
	pageSize := 1000
	opts := zoom.ListGroupMembersOptions{
		GroupID:  groupID,
		PageSize: &pageSize,
	}
	for {
		result, err := conn.ListGroupMembers(opts)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_group.listGroupMember", "query_error", err)
			return nil, err
		}
		for _, i := range result.Members {
			d.StreamListItem(ctx, i)
		}
		if result.NextPageToken == "" {
			break
		}
		opts.NextPageToken = &result.NextPageToken
	}
	return nil, nil
}
