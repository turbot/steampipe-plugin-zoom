package zoom

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableZoomAccountSettings(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_account_settings",
		Description: "Settings for the Zoom account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountSettings,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAccountSettings,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Hydrate: idString, Transform: transform.FromValue(), Description: "Account ID. Set to 'me' for the master account."},
			{Name: "schedule_meeting", Type: proto.ColumnType_JSON, Description: "Schedule meeting settings."},
			{Name: "in_meeting", Type: proto.ColumnType_JSON, Description: "In meeting settings."},
			{Name: "email_notification", Type: proto.ColumnType_JSON, Description: "Email notification settings."},
			{Name: "security", Type: proto.ColumnType_JSON, Description: "Security settings."},
			{Name: "recording", Type: proto.ColumnType_JSON, Description: "Recording settings."},
			{Name: "telephony", Type: proto.ColumnType_JSON, Description: "Telephony settings."},
			{Name: "tsp", Type: proto.ColumnType_JSON, Transform: transform.FromField("TSP"), Description: "TSP settings."},
			{Name: "integration", Type: proto.ColumnType_JSON, Description: "Integration settings."},
			{Name: "feature", Type: proto.ColumnType_JSON, Description: "Feature settings."},
			{Name: "meeting_authentication", Type: proto.ColumnType_JSON, Hydrate: getAccountSettingsMeetingAuthentication, Transform: transform.FromValue(), Description: "Meeting authentication options applied to the account."},
			{Name: "recording_authentication", Type: proto.ColumnType_JSON, Hydrate: getAccountSettingsRecordingAuthentication, Transform: transform.FromValue(), Description: "Recording authentication options applied to the account."},
			{Name: "meeting_security", Type: proto.ColumnType_JSON, Hydrate: getAccountSettingsMeetingSecurity, Description: "Meeting security settings applied to the account."},
			{Name: "trusted_domains", Type: proto.ColumnType_JSON, Hydrate: getAccountTrustedDomains, Description: "Associated domains allow all users with that email domain to be prompted to join the account."},
			{Name: "managed_domains", Type: proto.ColumnType_JSON, Hydrate: getAccountManagedDomains, Transform: transform.FromField("Domains"), Description: "Associated domains allow all users with that email domain to be prompted to join the account."},
		}),
	}
}

func listAccountSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	result, err := getAccountSettingsOption("", ctx, d, h)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}

func getAccountSettings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountSettingsOption("", ctx, d, h)
}

type authentication struct {
	Enabled               bool        `json:"enabled"`
	AuthenticationOptions interface{} `json:"authentication_options"`
}

func getAccountSettingsMeetingAuthentication(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	result, err := getAccountSettingsOption("meeting_authentication", ctx, d, h)
	if err != nil {
		return nil, err
	}
	settings := result.(zoom.AccountSettings)
	auth := authentication{
		Enabled:               settings.MeetingAuthentication,
		AuthenticationOptions: settings.AuthenticationOptions,
	}
	return auth, nil
}

func getAccountSettingsRecordingAuthentication(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	result, err := getAccountSettingsOption("recording_authentication", ctx, d, h)
	if err != nil {
		return nil, err
	}
	settings := result.(zoom.AccountSettings)
	auth := authentication{
		Enabled:               settings.RecordingAuthentication,
		AuthenticationOptions: settings.AuthenticationOptions,
	}
	return auth, nil
}

func getAccountSettingsMeetingSecurity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountSettingsOption("meeting_security", ctx, d, h)
}

func getAccountSettingsOption(option string, ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_account_settings.getAccountSettingsOption", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	if id == "" {
		id = "me"
	}
	opts := zoom.GetAccountSettingsOpts{
		AccountID: id,
	}
	if option != "" {
		opts.Option = option
	}
	result, err := conn.GetAccountSettings(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// Not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_account_settings.getAccountSettingsOption", "query_error", err)
		return nil, err
	}
	return result, nil
}

func getAccountTrustedDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_account_settings.getAccountTrustedDomainsOption", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	if id == "" {
		id = "me"
	}
	opts := zoom.GetAccountTrustedDomainsOpts{
		AccountID: id,
	}
	result, err := conn.GetAccountTrustedDomains(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// Not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_account_settings.getAccountTrustedDomainsOption", "query_error", err)
		return nil, err
	}
	return result, nil
}

func getAccountManagedDomains(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_account_settings.getAccountManagedDomains", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	if id == "" {
		id = "me"
	}
	opts := zoom.GetAccountManagedDomainsOpts{
		AccountID: id,
	}
	result, err := conn.GetAccountManagedDomains(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// Not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_account_settings.getAccountManagedDomains", "query_error", err)
		return nil, err
	}
	// Always return an array
	if result.Domains == nil {
		result.Domains = []string{}
	}
	return result, nil
}
