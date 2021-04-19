package zoom

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type zoomConfig struct {
	APIKey    *string `cty:"api_key"`
	APISecret *string `cty:"api_secret"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"api_secret": {
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
