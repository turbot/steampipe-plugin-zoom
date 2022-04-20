package zoom

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
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
		},
	}
}

type zoomAccountID struct {
	AccountID string
}

func getAccountID(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// This is called a lot and never changes for a connection, so cache it
	cacheKey := "zoom_account_id"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		plugin.Logger(ctx).Warn("zoom_user.getAccountID", "cache", "hit!")
		return cachedData.(zoomAccountID), nil
	}

	plugin.Logger(ctx).Warn("zoom_user.getAccountID", "cache", "miss!")

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_user.getAccountID", "connection_error", err)
		return nil, err
	}
	opts := zoom.GetUserOpts{
		EmailOrID: "me",
	}
	user, err := conn.GetUser(opts)
	if err != nil {
		plugin.Logger(ctx).Error("zoom_user.getAccountID", "query_error", err, "opts", opts)
		return nil, err
	}

	result := zoomAccountID{
		AccountID: user.AccountID,
	}

	plugin.Logger(ctx).Warn("zoom_user.getAccountID", "cache", "set!")

	// Save to cache
	setResult := d.ConnectionManager.Cache.Set(cacheKey, result)

	plugin.Logger(ctx).Warn("zoom_user.getAccountID", "setResult", setResult)
	plugin.Logger(ctx).Warn("zoom_user.getAccountID", "result", result)

	return result, nil
}
