package zoom

import (
	"context"

	"github.com/bigdatasourav/zoom-lib-golang"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Append the standard zoom account columns used by many tables
func zoomAccountColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonZoomAccountColumns()...)
}

// column definitions for the common columns
func commonZoomAccountColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getAccountID,
			Description: "Zoom account ID.",
			Transform:   transform.FromValue(),
		},
	}
}

func getAccountID(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := getAccountIDCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(string), nil
}

var getAccountIDCached = plugin.HydrateFunc(getAccountIDUncached).Memoize()

func getAccountIDUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	zoomConfig := GetConfig(d.Connection)
	var accountID string

	if zoomConfig.AccountID != nil {
		accountID = *zoomConfig.AccountID
	} else {
		conn, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.getAccountIDUncached", "connection_error", err)
			return nil, err
		}
		opts := zoom.GetUserOpts{
			EmailOrID: "me",
		}
		user, err := conn.GetUser(opts)
		if err != nil {
			plugin.Logger(ctx).Error("zoom_user.getAccountIDUncached", "query_error", err, "opts", opts)
			return nil, err
		}
		accountID = user.AccountID
	}

	return accountID, nil
}
