## trellogo
![build](https://github.com/utkuufuk/trellogo/workflows/trellogo/badge.svg?branch=master)
![Latest GitHub release](https://img.shields.io/github/release/utkuufuk/trellogo.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/utkuufuk/trellogo)](https://goreportcard.com/report/github.com/utkuufuk/trellogo)

A command line tool for interacting with the Trello API.

### Configuration
Copy and rename `config.example.yml` as `.trellogo.yml`, and put it into your `$HOME` directory. Then set your own values in `~/.trellogo.yml`.

### Usage
You can either download the latest binary from the [Releases](https://github.com/utkuufuk/trellogo/releases) page or build it yourself using the following command:
```sh
go build ./cmd/trellogo
```

| Command | Description |
|:-|:-|
| `trellogo list`               | lists cards from all lists included in config |
| `trellogo list <L1> <L2> ...` | lists cards from only the selected lists (as aliased in config) |

### References
 * [Trello REST API](https://developer.atlassian.com/cloud/trello/rest/)
 * [Trello REST API Help](http://www.trello.org/help.html)
