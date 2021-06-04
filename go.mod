module github.com/turbot/steampipe-plugin-zoom

go 1.16

replace github.com/himalayan-institute/zoom-lib-golang => github.com/e-gineer/zoom-lib-golang v1.0.1-0.20210430204608-85977a1550a0

require (
	github.com/golang/protobuf v1.4.3
	github.com/himalayan-institute/zoom-lib-golang v1.0.0
	github.com/turbot/steampipe-plugin-sdk v0.3.0-rc.0.0.20210526220805-6ca2e720464f
)
