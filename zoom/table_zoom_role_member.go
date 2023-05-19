package zoom

import (
	"context"

	"github.com/bigdatasourav/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableZoomRoleMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_role_member",
		Description: "Members of a Zoom role.",
		List: &plugin.ListConfig{
			Hydrate:    listRoleMember,
			KeyColumns: plugin.SingleColumn("role_id"),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "role_id", Type: proto.ColumnType_STRING, Hydrate: roleIDString, Transform: transform.FromValue(), Description: "Role ID."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID"), Description: "User ID of the member."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the member."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "First name of the member."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "Last name of the member."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the member."},
			{Name: "department", Type: proto.ColumnType_STRING, Description: "Department of the member."},
		}),
	}
}

func listRoleMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	roleID := d.EqualsQuals["role_id"].GetStringValue()
	pageSize := 1000
	opts := zoom.ListRoleMembersOptions{
		RoleID:   roleID,
		PageSize: &pageSize,
	}
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil { // check if JWT creds is set
		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_role.connect.listRoleMember", "connection_error", err)
			return nil, err
		}
		for {
			result, err := conn.ListRoleMembers(opts)
			if err != nil {
				plugin.Logger(ctx).Error("zoom_role.connect.listRoleMember", "query_error", err)
				return nil, err
			}
			for _, i := range result.Members {
				d.StreamListItem(ctx, i)
			}
			if result.NextPageToken == "" {
				break
			}
			opts.NextPageToken = result.NextPageToken
		}
	} else { // check if server-to-server oauth creds is set
		conn, err := connectOAuth(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_role.connectOAuth.listRoleMember", "connection_error", err)
			return nil, err
		}
		for {
			result, err := conn.ListRoleMembers(opts)
			if err != nil {
				plugin.Logger(ctx).Error("zoom_role.connectOAuth.listRoleMember", "query_error", err)
				return nil, err
			}
			for _, i := range result.Members {
				d.StreamListItem(ctx, i)
			}
			if result.NextPageToken == "" {
				break
			}
			opts.NextPageToken = result.NextPageToken
		}
	}
	return nil, nil
}
