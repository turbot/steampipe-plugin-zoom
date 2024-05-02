package zoom

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-zoom",
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "account_id",
				Hydrate: getAccountID,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"zoom_account_settings":      tableZoomAccountSettings(ctx),
			"zoom_account_lock_settings": tableZoomAccountLockSettings(ctx),
			"zoom_cloud_recording":       tableZoomCloudRecording(ctx),
			"zoom_group":                 tableZoomGroup(ctx),
			"zoom_group_member":          tableZoomGroupMember(ctx),
			"zoom_meeting":               tableZoomMeeting(ctx),
			"zoom_my_user":               tableZoomMyUser(ctx),
			"zoom_role":                  tableZoomRole(ctx),
			"zoom_role_member":           tableZoomRoleMember(ctx),
			"zoom_user":                  tableZoomUser(ctx),
		},
	}
	return p
}
