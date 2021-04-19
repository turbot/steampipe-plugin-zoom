![image](https://hub.steampipe.io/images/plugins/turbot/zoom-social-graphic.png)

# Zoom Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from Zoom.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/zoom)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/zoom/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-zoom/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install zoom
```

Run a query:
```sql
select
  email,
  personal_meeting_url
from
  zoom_user
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone git@github.com:turbot/steampipe-plugin-zoom
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
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [MPL-2.0 open source license](https://github.com/turbot/steampipe-plugin-zoom/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Zoom Plugin](https://github.com/turbot/steampipe-plugin-zoom/labels/help%20wanted)
