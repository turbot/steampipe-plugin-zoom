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

List your Zoom users:

```sql
select
  email,
  personal_meeting_url
from
  zoom_user;
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

| Item        | Description                                                                                                                                                                                                                                                  |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Credentials | [Create a Server-to-Server OAuth app](https://developers.zoom.us/docs/internal-apps/create/) to get the Account ID, Client ID and Client Secret or [create an SDK App](https://marketplace.zoom.us/docs/guides/build/sdk-app) to get the API key and secret. |
| Permissions | Server-to-Server OAuth apps or SDK / JWT apps can access all Zoom APIs.                                                                                                                                                                                      |
| Radius      | Each connection represents a single Zoom account.                                                                                                                                                                                                            |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/zoom.spc`)<br />2. Credentials specified in environment variables, e.g., `ZOOM_ACCOUNT_ID`, `ZOOM_CLIENT_ID` and `ZOOM_CLIENT_SECRET`.                                        |

### Configuration

Installing the latest zoom plugin will create a config file (`~/.steampipe/config/zoom.spc`) with a single connection named `zoom`:

Configure your account details in `~/.steampipe/config/zoom.spc`:

```hcl
connection "zoom" {
  plugin = "zoom"

  # Zoom API credentials are available to users with Developer role in the account.
  # You need to create a Server-to-Server OAuth app(https://developers.zoom.us/docs/internal-apps/create) or a SDK/JWT APP (https://marketplace.zoom.us/docs/guides/build/sdk-app)to get the credentials.
  # It is recommended that you create Server-to-Server OAuth as JWT app is deprecated On June 1, 2023 and will be disabled on September 1, 2023. https://developers.zoom.us/docs/internal-apps/jwt-faq/

  ## Server-to-Server OAuth app credentials

  # `account_id`(required) - The Zoom account ID.
  # Can also be set with the ZOOM_ACCOUNT_ID environment variable.
  # account_id    = "Xt1aUD4WQ56w7hDhVbtDp"

  # `client_id`(required) - The Zoom Client ID provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_ID environment variable.
  # client_id    = "MZw2piRfTsOdpwx2Dh5U"

  # `client_secret`(required) - The Zoom Client Secret provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_SECRET environment variable.
  # client_secret    = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"

  ## SDK/JWT app credentials

  # `api_key`(required)   = The Zoom API key provided by SDK/JWT OAuth app.
  # Can also be set with the ZOOM_API_KEY environment variable.
  # api_key    = "LFMU3oagTjO8_5sYKQVe"

  # `api_secret`(required) = The Zoom API secret provided by SDK/JWT OAuth app.
  # Can also be set with the ZOOM_API_SECRET environment variable.
  # api_secret = "PKS96L69nWSFK2y0A07R2k7xGryVbcWiem"
}
```

## Configuring Zoom Credentials

You may specify the AccountID, ClientID and ClientSecret to authenticate:

- `account_id`: The Zoom account ID.
- `client_id`: The Zoom Client ID provided by Server-to-Server OAuth app.
- `client_secret`: The Zoom Client Secret provided by Server-to-Server OAuth app.

```hcl
connection "zoom" {
  plugin        = "zoom"
  account_id    = "Xt1aUD4WQ56w7hDhVbtDp"
  client_id     = "MZw2piRfTsOdpwx2Dh5U"
  client_secret = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"
}
```

or you may specify the APIKey and APISecret to authenticate:

- `api_key`: The Zoom API key provided by SDK/JWT OAuth app.
- `api_secret`: The Zoom API secret provided by SDK/JWT OAuth app.

```hcl
connection "zoom" {
  plugin     = "zoom"
  api_key    = "9m_kAcfuTlW_JCrvoMYK6g"
  api_secret = "lEEDVf3SgyQWckN3ASqMpXWpCixkwMzgnZY7"
}
```

or through environment variables

The Zoom plugin will use the Zoom environment variable to obtain credentials **only if the `account_id`,`client_id`, and`client_secret` or `api_key` and `api_secret` is not specified** in the connection:

```sh
export ZOOM_ACCOUNT_ID="Xt1aUD4WQ56w7hDhVbtDp"
export ZOOM_CLIENT_ID="MZw2piRfTsOdpwx2Dh5U"
export ZOOM_CLIENT_SECRET="04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"
export ZOOM_API_KEY="9m_kAcfuTlW_JCrvoMYK6g"
export ZOOM_API_SECRET="lEEDVf3SgyQWckN3ASqMpXWpCixkwMzgnZY7"
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-zoom
- Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
