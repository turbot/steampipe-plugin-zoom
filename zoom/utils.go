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
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*zoom.Client), nil
}

func connectOAuth(ctx context.Context, d *plugin.QueryData) (*zoom.OAuthClient, error) {
	conn, err := connectOAuthCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*zoom.OAuthClient), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()
var connectOAuthCached = plugin.HydrateFunc(connectOAuthUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	// Default to using env vars (#2)
	apiKey := os.Getenv("ZOOM_API_KEY")
	apiSecret := os.Getenv("ZOOM_API_SECRET")

	// But prefer the config (#1)
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil {
		apiKey = *zoomConfig.APIKey
	}
	if zoomConfig.APISecret != nil {
		apiSecret = *zoomConfig.APISecret
	}

	if apiKey == "" || apiSecret == "" {
		// Credentials not set
		return nil, errors.New("api_key and api_secret must be configured")
	}

	// Configure to automatically wait 1 sec between requests, per Zoom API requirements
	conn := zoom.NewClient(apiKey, apiSecret)

	return conn, nil
}

func connectOAuthUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
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
	conn := zoom.NewOAuthClient(accountID, clientID, clientSecret)

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
