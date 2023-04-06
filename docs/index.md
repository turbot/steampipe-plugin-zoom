---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/zoom.svg"
brand_color: "#2D8CFF"
display_name: Zoom
name: zoom
description: Steampipe plugin for querying Zoom meetings, webinars, users and more.
og_description: Query Zoom with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/zoom-social-graphic.png"
---

# Zoom + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Zoom](https://zoom.us) provides videotelephony and online chat services through a cloud-based peer-to-peer software platform and is used for teleconferencing, telecommuting, distance education, and social relations.

For example:
```sql
select
  email,
  personal_meeting_url
from
  zoom_user
```

```
+--------------------------+--------------------------------------------------------------------------+
| email                    | personal_meeting_url                                                     |
+--------------------------+--------------------------------------------------------------------------+
| jim@dundermifflin.com    | https://turbot.zoom.us/j/9694476416                                      |
| dwight@dundermifflin.com | https://turbot.zoom.us/j/1453171280?pwd=bWloMG5Ic0JrKFQ2SGJaUmNxZXhNQT09 |
+--------------------------+--------------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/zoom/tables)**

## Get started

### Install

Download and install the latest Zoom plugin:

```bash
steampipe plugin install zoom
```

### Credentials

| Item | Description |
| - | - |
| Credentials | [Sign in](https://zoom.us/signin#/login) to Zoom to generate credentials. [Create a JWT App](https://marketplace.zoom.us/develop/create) to get the API key and secret. |
| Permissions | JWT apps can access all Zoom APIs. |
| Radius | Each connection represents a single Zoom account. |
| Resolution |  1. `api_key` and `api_secret` in Steampipe config.<br />2. `ZOOM_API_KEY` and `ZOOM_API_SECRET` environment variables. |

### Configuration

Installing the latest zoom plugin will create a config file (`~/.steampipe/config/zoom.spc`) with a single connection named `zoom`:

```hcl
connection "zoom" {
  plugin     = "zoom"
  api_key    = "9m_kAcfuTlW_JCrvoMYK6g"
  api_secret = "lEEDVf3SgyQWckN3ASqMpXWpCixkwMzgnZY7"
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-zoom
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
