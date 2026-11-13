package zoom

import (
	"context"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/zoom-lib-golang"
)

func shouldRetryError(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
	if e, ok := err.(*zoom.APIError); ok && e.Code == 429 {
		log.Println("[WARN] Rate Limit Error", e)
		return true
	}
	return false
}
