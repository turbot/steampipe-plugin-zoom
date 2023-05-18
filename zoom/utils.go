package zoom

import (
	"context"
	"errors"
	"os"

	"github.com/bigdatasourav/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*zoom.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "zoom"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*zoom.Client), nil
	}

	// Default to using env vars (#2)
	accountID := os.Getenv("ZOOM_ACCOUNT_ID")
	clientID := os.Getenv("ZOOM_CLIENT_ID")
	clientSecret := os.Getenv("ZOOM_CLIENT_SECRET")

	// But prefer the config (#1)
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.AccountID != nil {
		accountID = *zoomConfig.AccountID
	}
	if zoomConfig.ClientID != nil {
		clientID = *zoomConfig.ClientID
	}
	if zoomConfig.ClientSecret != nil {
		clientSecret = *zoomConfig.ClientSecret
	}

	if accountID == "" || clientID == "" || clientSecret == "" {
		// Credentials not set
		return nil, errors.New("account_id, client_id and client_secret must be configured")
	}

	// Configure to automatically wait 1 sec between requests, per Zoom API requirements
	conn := zoom.NewClient(accountID, clientID, clientSecret)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func timeToTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	switch t := d.Value.(type) {
	case zoom.Time:
		return t.String(), nil
	case *zoom.Time:
		if t != nil {
			return (*t).String(), nil
		}
	}
	return nil, nil
}

func idString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	// Default to the master account, called me
	if id == "" {
		id = "me"
	}
	return id, nil
}

func roleIDString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	id := quals["role_id"].GetStringValue()
	return id, nil
}

func groupIDString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	id := quals["group_id"].GetStringValue()
	return id, nil
}
