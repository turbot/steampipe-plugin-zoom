package zoom

import (
	"log"

	"github.com/himalayan-institute/zoom-lib-golang"
)

func shouldRetryError(err error) bool {
	if e, ok := err.(*zoom.APIError); ok && e.Code == 429 {
		log.Println("[WARN] Rate Limit Error", e)
		return true
	}
	return false
}
