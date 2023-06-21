package zoom

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type zoomConfig struct {
	AccountID    *string `cty:"account_id"`
	APIKey       *string `cty:"api_key"`
	APISecret    *string `cty:"api_secret"`
	ClientID     *string `cty:"client_id"`
	ClientSecret *string `cty:"client_secret"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"account_id": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
	"api_secret": {
		Type: schema.TypeString,
	},
	"client_id": {
		Type: schema.TypeString,
	},
	"client_secret": {
		Type: schema.TypeString,
	},
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
