package zoom

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type zoomConfig struct {
	AccountID    *string `hcl:"account_id"`
	APIKey       *string `hcl:"api_key"`
	APISecret    *string `hcl:"api_secret"`
	ClientID     *string `hcl:"client_id"`
	ClientSecret *string `hcl:"client_secret"`
}

func ConfigInstance() interface{} {
	return &zoomConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) zoomConfig {
	if connection == nil || connection.Config == nil {
		return zoomConfig{}
	}
	config, _ := connection.Config.(zoomConfig)
	return config
}
