## trellogo
![build](https://github.com/utkuufuk/trellogo/workflows/trellogo/badge.svg?branch=master)

A command line tool for interacting with the Trello API.

### Configuration
Copy and rename `config.example.yml` as `config.yml`, then set your own values in `config.yml`.

### Usage
| Command | Description |
|:-|:-|
| `go run ./cmd list`               | lists cards from all lists included in config |
| `go run ./cmd list <L1> <L2> ...` | lists cards from only the selected lists (as aliased in config) |

### References
 * [Trello REST API](https://developer.atlassian.com/cloud/trello/rest/)
 * [Trello REST API Help](http://www.trello.org/help.html)
