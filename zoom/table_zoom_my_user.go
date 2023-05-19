package zoom

import (
	"context"

	"github.com/bigdatasourav/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableZoomMyUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_my_user",
		Description: "Basic information about the authenticated user being used by the plugin.",
		List: &plugin.ListConfig{
			Hydrate: listMyUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "User ID."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "User's email address."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "User's first name."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "User's last name."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "Zoom account ID that contains this user."},
		},
	}
}

func listMyUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	opts := zoom.GetUserOpts{
		EmailOrID: "me",
	}
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil { // check if JWT creds is set
		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.connect.listMyUser", "connection_error", err)
			return nil, err
		}
		result, err := conn.GetUser(opts)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.connect.listMyUser", "query_error", err, "opts", opts)
			return nil, err
		}
		d.StreamListItem(ctx, result)
	} else { // check if server-to-server oauth creds is set
		conn, err := connectOAuth(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.connectOAuth.listMyUser", "connection_error", err)
			return nil, err
		}
		result, err := conn.GetUser(opts)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.connectOAuth.listMyUser", "query_error", err, "opts", opts)
			return nil, err
		}
		d.StreamListItem(ctx, result)
	}
	return nil, nil
}
