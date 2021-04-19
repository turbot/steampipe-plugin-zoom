module github.com/turbot/steampipe-plugin-zoom

go 1.16

replace github.com/himalayan-institute/zoom-lib-golang => github.com/e-gineer/zoom-lib-golang v1.0.1-0.20210419004004-b956a344cc40

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/himalayan-institute/zoom-lib-golang v1.0.0
	github.com/turbot/steampipe-plugin-sdk v0.2.7
)
