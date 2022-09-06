package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-zoom/zoom"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: zoom.Plugin})
}
