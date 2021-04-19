package zoom

import (
	"context"
	"errors"
	"os"

	"github.com/himalayan-institute/zoom-lib-golang"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*zoom.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "zoom"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*zoom.Client), nil
	}

	// Default to using env vars (#2)
	apiKey := os.Getenv("ZOOM_API_KEY")
	apiSecret := os.Getenv("ZOOM_API_SECRET")

	// But prefer the config (#1)
	zoomConfig := GetConfig(d.Connection)
	if &zoomConfig != nil {
		if zoomConfig.APIKey != nil {
			apiKey = *zoomConfig.APIKey
		}
		if zoomConfig.APISecret != nil {
			apiSecret = *zoomConfig.APISecret
		}
	}

	if apiKey == "" || apiSecret == "" {
		// Credentials not set
		return nil, errors.New("api_key and api_secret must be configured")
	}

	// Configure to automatically wait 1 sec between requests, per Zoom API requirements
	conn := zoom.NewClient(apiKey, apiSecret)

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
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	// Default to the master account, called me
	if id == "" {
		id = "me"
	}
	return id, nil
}

func roleIDString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	id := quals["role_id"].GetStringValue()
	return id, nil
}

func groupIDString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	id := quals["group_id"].GetStringValue()
	return id, nil
}
