package zoom

import (
	"context"

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

	// This is called a lot and never changes for a connection, so cache it
	cacheKey := "zoom_account_id"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*string), nil
	}

	zoomConfig := GetConfig(d.Connection)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, zoomConfig.AccountID)

	return zoomConfig.AccountID, nil
}
