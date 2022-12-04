# jsonfmt

[![GoDoc](https://pkg.go.dev/badge/github.com/ansidev/jsonfmt?status.svg)](https://pkg.go.dev/github.com/ansidev/jsonfmt?tab=doc)
[![Release](https://img.shields.io/github/release/ansidev/jsonfmt.svg)](https://github.com/ansidev/jsonfmt/releases)
[![Build Status](https://github.com/ansidev/jsonfmt/workflows/main/badge.svg?branch=main)](https://github.com/ansidev/jsonfmt/actions?query=branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/ansidev/jsonfmt)](https://goreportcard.com/report/github.com/ansidev/jsonfmt)
[![Sourcegraph](https://sourcegraph.com/github.com/ansidev/jsonfmt/-/badge.svg)](https://sourcegraph.com/github.com/ansidev/jsonfmt?badge)

## Getting Started
### Installation

```
go install github.com/ansidev/jsonfmt@latest
```

### Usage

```
NAME:
   jsonfmt - JSON Formatter CLI

USAGE:
   jsonfmt [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --color, -c               Enable color, only applied if the output is terminal. Default is false. (default: false)
   --debug, -d               Enable debug mode. Default is false. (default: false)
   --help, -h                show help (default: false)
   --indent value, -i value  Indent config. Format: '{indent_style}:{indent_value}'.
                             Indent style: w (white space), t (tab), s (string).
                             If indent_style is w or t, the indent value will be a number (indent size).
                             If indent_style is s, the indent value will be indent string.
                             Default indentation is 2 spaces. (default: "w:2")
   --minify, -m              Minify formatted json. These options will be ignored: -i, -w. Default is false. (default: false)
   --output value, -o value  Output path. Use this option if you want to save to file, by default the result will be printed to the stdout.
   --prefix value, -p value  Prefix is a prefix for all lines. Default is an empty string.
   --sort-keys, -s           Sort keys will sort the keys alphabetically. Default is false. (default: false)
   --width value, -w value   Width is a max column width for single line arrays. Default is 80. (default: 80)
```

## Contact

Le Minh Tri [@ansidev](https://ansidev.xyz/about).

## License

`jsonfmt` source code is available under the [MIT License](/LICENSE).