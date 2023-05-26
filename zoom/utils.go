package zoom

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/zoom-lib-golang"

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

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	// Default to using env vars (#2)
	apiKey := os.Getenv("ZOOM_API_KEY")
	apiSecret := os.Getenv("ZOOM_API_SECRET")
	accountID := os.Getenv("ZOOM_ACCOUNT_ID")
	clientID := os.Getenv("ZOOM_CLIENT_ID")
	clientSecret := os.Getenv("ZOOM_CLIENT_SECRET")

	// But prefer the config (#1)
	zoomConfig := GetConfig(d.Connection)
	if zoomConfig.APIKey != nil {
		apiKey = *zoomConfig.APIKey
	}
	if zoomConfig.APISecret != nil {
		apiSecret = *zoomConfig.APISecret
	}
	if zoomConfig.AccountID != nil {
		accountID = *zoomConfig.AccountID
	}
	if zoomConfig.ClientID != nil {
		clientID = *zoomConfig.ClientID
	}
	if zoomConfig.ClientSecret != nil {
		clientSecret = *zoomConfig.ClientSecret
	}

	if (accountID == "" || clientID == "" || clientSecret == "") && (apiKey == "" || apiSecret == "") {
		// Credentials not set
		return nil, errors.New("Server-to-Server oauth app or JWT app credentials must be configured")
	}

	// prefer server-to-server oauth app creds
	if accountID != "" || clientID != "" || clientSecret != "" {
		conn := zoom.NewClient("", "", accountID, clientID, clientSecret)
		return conn, nil
	} else {
		conn := zoom.NewClient(apiKey, apiSecret, "", "", "")
		return conn, nil
	}
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
