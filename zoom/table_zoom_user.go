package zoom

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableZoomUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "zoom_user",
		Description: "A Zoom account can have one or more users.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getUser,
			KeyColumns: plugin.AnyColumn([]string{"id", "email"}),
		},
		Columns: zoomAccountColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "User ID."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "User's first name."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "User's last name."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "User's email address."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "User's status: pending, active or inactive."},
			// Other Columns
			{Name: "cms_user_id", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "CMS ID of user, only enabled for Kaltura integration."},
			{Name: "company", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "User's company."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToTimestamp), Description: "The time when user's account was created."},
			{Name: "custom_attributes", Type: proto.ColumnType_JSON, Description: "Custom attributes, if any are assigned."},
			{Name: "dept", Type: proto.ColumnType_STRING, Description: "Department, if provided by the user."},
			{Name: "group_ids", Type: proto.ColumnType_JSON, Description: "IDs of groups where the user is a member."},
			{Name: "host_key", Type: proto.ColumnType_STRING, Description: "The host key of the user."},
			{Name: "im_group_ids", Type: proto.ColumnType_JSON, Description: "IDs of IM directory groups where the user is a member."},
			{Name: "jid", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "User's JID."},
			{Name: "job_title", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "User's job title."},
			{Name: "language", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "Default language for the Zoom Web Portal."},
			{Name: "last_client_version", Type: proto.ColumnType_STRING, Description: "The last client version that user used to login."},
			{Name: "last_login_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastLoginTime").Transform(timeToTimestamp), Description: "User's last login time. There is a three-days buffer period for this field. For example, if user first logged in on 2020-01-01 and then logged out and logged in on 2020-01-02, the value of this field will still reflect the login time of 2020-01-01. However, if the user logs in on 2020-01-04, the value of this field will reflect the corresponding login time since it exceeds the three-day buffer period."},
			{Name: "location", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "User's location."},
			{Name: "login_type", Type: proto.ColumnType_INT, Hydrate: getUser, Description: "Login type. 0 - Facebook, 1 - Google, 99 - API, 100 - ZOOM, 101 - SSO"},
			{Name: "personal_meeting_url", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "User's personal meeting url."},
			{Name: "phone_numbers", Type: proto.ColumnType_JSON, Hydrate: getUser, Description: "User phone number, including verification status."},
			{Name: "pic_url", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "The URL for user's profile picture."},
			{Name: "plan_united_type", Type: proto.ColumnType_STRING, Description: "This field is returned if the user is enrolled in the Zoom United plan."},
			{Name: "pmi", Type: proto.ColumnType_INT, Transform: transform.FromField("PMI"), Description: "Personal meeting ID of the user."},
			{Name: "role_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the role assigned to the user."},
			{Name: "timezone", Type: proto.ColumnType_STRING, Description: "The time zone of the user."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "User's plan type: 1 - Basic, 2 - Licensed or 3 - On-prem."},
			{Name: "use_pmi", Type: proto.ColumnType_BOOL, Hydrate: getUser, Description: "Use Personal Meeting ID for instant meetings."},
			{Name: "vanity_url", Type: proto.ColumnType_STRING, Hydrate: getUser, Description: "Personal meeting room URL, if the user has one."},
			{Name: "verified", Type: proto.ColumnType_INT, Transform: transform.FromField("Verified"), Description: "Display whether the user's email address for the Zoom account is verified or not. 1 - Verified user email. 0 - User's email not verified."},
		}),
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_user.listUser", "connection_error", err)
		return nil, err
	}

	opts := zoom.ListUsersOptions{
		PageSize:      300,
		IncludeFields: &[]string{"custom_attributes", "host_key"},
	}
	for {
		result, err := conn.ListUsers(opts)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.listUser", "query_error", err)
			return nil, err
		}
		for _, i := range result.Users {
			d.StreamListItem(ctx, i)
		}
		if result.NextPageToken == "" {
			break
		}
		opts.NextPageToken = &result.NextPageToken
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_user.getUser", "connection_error", err)
		return nil, err
	}
	// Work with quals: email, then id
	quals := d.KeyColumnQuals
	emailOrID := quals["email"].GetStringValue()
	id := quals["id"].GetStringValue()
	if emailOrID == "" {
		emailOrID = id
	}
	// Prefer the hydration from the list if available
	if user, ok := h.Item.(zoom.User); ok {
		emailOrID = user.ID
	}
	opts := zoom.GetUserOpts{
		EmailOrID: emailOrID,
	}
	result, err := conn.GetUser(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// User not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_user.getUser", "query_error", err)
		return nil, err
	}
	return result, nil
}

func getUserPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_user.getUserPermissions", "connection_error", err)
		return nil, err
	}
	user := h.Item.(zoom.User)
	opts := zoom.GetUserPermissionsOpts{
		UserID: user.ID,
	}
	result, err := conn.GetUserPermissions(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			// User not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("zoom_user.getUserPermissions", "query_error", err)
		return nil, err
	}
	return result, nil
}
