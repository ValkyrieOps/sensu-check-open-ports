[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/ValkyrieOps/sensu-check-open-ports)
![Go Test](https://github.com/ValkyrieOps/sensu-check-open-ports/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/ValkyrieOps/sensu-check-open-ports/workflows/goreleaser/badge.svg)

# sensu-check-open-ports

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The sensu-check-open-ports is a [Sensu Check][6] that determines open port count per user

## Files

## Usage examples

```
Sensu Check for Open Ports

Usage:
  sensu-check-open-ports [flags]
  sensu-check-open-ports [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --crit int      Critical threshold - count of open ports required for critical state
  -h, --help          help for sensu-check-open-ports
  -u, --user string   User to query for open port count (default "sensu")
  -w, --warn int      Warning threshold - count of open ports required for warning state

Use "sensu-check-open-ports [command] --help" for more information about a command.

```
## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add ValkyrieOps/sensu-check-open-ports
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/ValkyrieOps/sensu-check-open-ports].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-check-open-ports
  namespace: default
spec:
  command: sensu-check-open-ports --example example_arg
  subscriptions:
  - system
  runtime_assets:
  - ValkyrieOps/sensu-check-open-ports
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the sensu-check-open-ports repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/check-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/check-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu-community/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
