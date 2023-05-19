package zoom

import (
	"context"

	"github.com/bigdatasourav/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableZoomAccountLockSettings(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_account_lock_settings",
		Description: "Lock settings for the Zoom account. If true, the setting is locked at the account level.",
		List: &plugin.ListConfig{
			Hydrate: listAccountLockSettings,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAccountLockSettings,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Hydrate: idString, Transform: transform.FromValue(), Description: "Account ID. Set to 'me' for the master account."},
			{Name: "schedule_meeting", Type: proto.ColumnType_JSON, Description: "Schedule meeting lock settings."},
			{Name: "in_meeting", Type: proto.ColumnType_JSON, Description: "In meeting lock settings."},
			{Name: "email_notification", Type: proto.ColumnType_JSON, Description: "Email notification lock settings."},
			{Name: "recording", Type: proto.ColumnType_JSON, Description: "Recording lock settings."},
			{Name: "telephony", Type: proto.ColumnType_JSON, Description: "Telephony lock settings."},
			{Name: "tsp", Type: proto.ColumnType_JSON, Description: "TSP lock settings."},
			{Name: "meeting_security", Type: proto.ColumnType_JSON, Hydrate: getAccountLockSettingsMeetingSecurity, Description: "Meeting security lock settings."},
		}),
	}
}

func listAccountLockSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	result, err := getAccountLockSettingsOption("", ctx, d, h)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}

func getAccountLockSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountLockSettingsOption("", ctx, d, h)
}

func getAccountLockSettingsMeetingSecurity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountLockSettingsOption("meeting_security", ctx, d, h)
}

func getAccountLockSettingsOption(option string, ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()
	if id == "" {
		id = "me"
	}
	opts := zoom.GetAccountLockSettingsOpts{
		AccountID: id,
	}
	if option != "" {
		opts.Option = option
	}

	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil { // check if JWT creds is set
		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_account_lock_settings.connect.getAccountLockSettingsOption", "connection_error", err)
			return nil, err
		}
		result, err := conn.GetAccountLockSettings(opts)
		if err != nil {
			if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
				// Not found
				return nil, nil
			}
			plugin.Logger(ctx).Error("zoom_account_lock_settings.connect.getAccountLockSettingsOption", "query_error", err)
			return nil, err
		}
		return result, nil
	} else { // check if server-to-server oauth creds is set
		conn, err := connectOAuth(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_account_lock_settings.connectOAuth.getAccountLockSettingsOption", "connection_error", err)
			return nil, err
		}

		result, err := conn.GetAccountLockSettings(opts)
		if err != nil {
			if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
				// Not found
				return nil, nil
			}
			plugin.Logger(ctx).Error("zoom_account_lock_settings.connectOAuth.getAccountLockSettingsOption", "query_error", err)
			return nil, err
		}
		return result, nil
	}
}
