![image](https://hub.steampipe.io/images/plugins/turbot/zoom-social-graphic.png)

# Zoom Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from Zoom.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/zoom)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/zoom/tables)
- Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-zoom/issues)

## Quick start

### Install

Download and install the latest Zoom plugin:

```bash
steampipe plugin install zoom
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/zoom#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/zoom#configuration).

Configure your account details in `~/.steampipe/config/zoom.spc`:

You may specify the account ID, client ID, and client secret to authenticate:

```hcl
connection "zoom" {
  plugin        = "zoom"
  account_id    = "Xt1aUD4WQ56w7hDhVbtDp"
  client_id     = "MZw2piRfTsOdpwx2Dh5U"
  client_secret = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"
}
```

Or through environment variables:

```sh
export ZOOM_ACCOUNT_ID="Xt1aUD4WQ56w7hDhVbtDp"
export ZOOM_CLIENT_ID="MZw2piRfTsOdpwx2Dh5U"
export ZOOM_CLIENT_SECRET="04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"
```

You may also use the deprecated [JWT App](https://hub.steampipe.io/plugins/turbot/zoom#jwt-application) mode of authentication.

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-zoom.git
cd steampipe-plugin-zoom
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/zoom.spc
```

Try it!

```
steampipe query
> .inspect zoom
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-zoom/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-zoom/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Zoom Plugin](https://github.com/turbot/steampipe-plugin-zoom/labels/help%20wanted)
