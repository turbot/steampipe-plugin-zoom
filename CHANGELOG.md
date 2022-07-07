## v0.3.1 [2022-07-08]

_Enhancements_

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
