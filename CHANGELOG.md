## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Bug fixes_

- Fixed the `zoom_account_settings` table to correctly return data instead of an error. ([#60](https://github.com/turbot/steampipe-plugin-zoom/pull/60))

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#61](https://github.com/turbot/steampipe-plugin-zoom/pull/61))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#61](https://github.com/turbot/steampipe-plugin-zoom/pull/61))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#56](https://github.com/turbot/steampipe-plugin-zoom/pull/56))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#56](https://github.com/turbot/steampipe-plugin-zoom/pull/56))

## v0.10.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#51](https://github.com/turbot/steampipe-plugin-zoom/pull/51))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#51](https://github.com/turbot/steampipe-plugin-zoom/pull/51))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-zoom/blob/main/docs/LICENSE). ([#51](https://github.com/turbot/steampipe-plugin-zoom/pull/51))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to  column `_ctx`, and fixing connection and potential divide-by-zero bugs. ([#50](https://github.com/turbot/steampipe-plugin-zoom/pull/50))

## v0.9.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#43](https://github.com/turbot/steampipe-plugin-zoom/pull/43))

## v0.9.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#41](https://github.com/turbot/steampipe-plugin-zoom/pull/41))
- Recompiled plugin with Go version `1.21`. ([#41](https://github.com/turbot/steampipe-plugin-zoom/pull/41))

## v0.8.0 [2023-06-21]

_What's new?_

- Added support for [Server-to-Server OAuth](https://developers.zoom.us/docs/internal-apps/s2s-oauth/#enable-the-server-to-server-oauth-role) authentication mechanism. Please refer the [Configuration](https://hub.steampipe.io/plugins/turbot/zoom#configuration) section for additional information. ([#35](https://github.com/turbot/steampipe-plugin-zoom/pull/35))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. ([#36](https://github.com/turbot/steampipe-plugin-zoom/pull/36))

## v0.7.0 [2023-05-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05) which fixes increased plugin initialization time due to multiple connections causing the schema to be loaded repeatedly. ([#33](https://github.com/turbot/steampipe-plugin-zoom/pull/33))

## v0.6.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#29](https://github.com/turbot/steampipe-plugin-zoom/pull/29))

## v0.5.1 [2023-02-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.12](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v4112-2023-02-09) which fixes the query caching functionality. ([#26](https://github.com/turbot/steampipe-plugin-zoom/pull/26))

## v0.5.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#23](https://github.com/turbot/steampipe-plugin-zoom/pull/23))
- Recompiled plugin with Go version `1.19`. ([#23](https://github.com/turbot/steampipe-plugin-zoom/pull/23))

## v0.4.0 [2022-07-21]

_Bug fixes_

- Fixed the `GetConfig` max concurrency configuration in the `zoom_meeting` table to resolve the plugin validation errors. ([#21](https://github.com/turbot/steampipe-plugin-zoom/pull/21))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#21](https://github.com/turbot/steampipe-plugin-zoom/pull/21))

## v0.3.1 [2022-07-08]

_Bug fixes_

- Reverted [steampipe-plugin-sdk v3.1.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v331--2022-06-30) update and recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) due to plugin incompatibilities with the new SDK version. ([#19](https://github.com/turbot/steampipe-plugin-zoom/pull/19))

## v0.3.0 [2022-07-07]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v331--2022-06-30) which includes several caching fixes. ([#17](https://github.com/turbot/steampipe-plugin-zoom/pull/17))

## v0.2.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#15](https://github.com/turbot/steampipe-plugin-zoom/pull/15))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#14](https://github.com/turbot/steampipe-plugin-zoom/pull/14))

## v0.1.0 [2021-12-08]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#11](https://github.com/turbot/steampipe-plugin-zoom/pull/11))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#10](https://github.com/turbot/steampipe-plugin-zoom/pull/10))

## v0.0.4 [2021-09-22]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#5](https://github.com/turbot/steampipe-plugin-zoom/pull/5))
- Changed plugin license to Apache 2.0 per [turbot/steampipe](https://github.com/turbot/steampipe/issues/488) ([#3](https://github.com/turbot/steampipe-plugin-zoom/pull/3))

## v0.0.3 [2021-06-04]

_What's new?_

- New tables added
  - [zoom_my_user](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_my_user)

_Enhancements_

- Add common column of `account_id` to all tables.

## v0.0.2 [2021-05-06]

_Enhancements_

- Use Steampipe SDK v0.2.8 with retry support ([#1](https://github.com/turbot/steampipe-plugin-zoom/pull/1))

## v0.0.1 [2021-04-30]

_What's new?_

- New tables added
  - [zoom_account_lock_settings](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_account_lock_settings)
  - [zoom_account_settings](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_account_settings)
  - [zoom_cloud_recording](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_cloud_recording)
  - [zoom_group](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_group)
  - [zoom_group_member](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_group_member)
  - [zoom_meeting](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_meeting)
  - [zoom_role](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_role)
  - [zoom_role_member](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_role_member)
  - [zoom_user](https://hub.steampipe.io/plugins/turbot/zoom/tables/zoom_user)
