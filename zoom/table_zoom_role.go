package zoom

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableZoomRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_role",
		Description: "A Zoom account can have one or more roles.",
		List: &plugin.ListConfig{
			Hydrate: listRole,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getRole,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Role ID."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Role name."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Role description."},
			{Name: "total_members", Type: proto.ColumnType_INT, Transform: transform.FromField("TotalMembers"), Description: "Total number of members in the role."},
			{Name: "privileges", Type: proto.ColumnType_JSON, Hydrate: getRole, Description: "Privileges assigned to the role."},
			{Name: "sub_account_privileges", Type: proto.ColumnType_JSON, Hydrate: getRole, Description: "Privileges for management of sub-accounts."},
		}),
	}
}

func listRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_role.listRole", "connection_error", err)
		return nil, err
	}
	result, err := conn.ListRoles()
	if err != nil {
		plugin.Logger(ctx).Error("zoom_role.listRole", "query_error", err)
		return nil, err
	}
	for _, i := range result.Roles {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_role.getRole", "connection_error", err)
		return nil, err
	}
	var id string
	if h.Item != nil {
		id = h.Item.(zoom.Role).ID
	} else {
		id = d.EqualsQuals["id"].GetStringValue()
	}
	opts := zoom.GetRoleOpts{
		ID: id,
	}
	result, err := conn.GetRole(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// Role not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_role.getRole", "query_error", err)
		return nil, err
	}
	return result, nil
}
