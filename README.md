# godestone

[![Documentation Badge](https://img.shields.io/badge/docs-pkg.go.dev-007D9C)](https://pkg.go.dev/github.com/xivapi/godestone/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/xivapi/godestone)](https://goreportcard.com/report/github.com/xivapi/godestone)
[![Dependencies](https://img.shields.io/librariesio/github/xivapi/godestone)](https://libraries.io/github/xivapi/godestone)
[![Code Size](https://img.shields.io/github/languages/code-size/xivapi/godestone)](https://github.com/xivapi/godestone)
[![Repo Size](https://img.shields.io/github/repo-size/xivapi/godestone)](https://github.com/xivapi/godestone)

Go scraper for The Lodestone.

## Installation
Just add the package to your `go.mod` or run `go get github.com/xivapi/godestone/v2`.

Also be sure to install a data provider service to initialize the scraper.

## Data providers
Package|Size|Description
---|---|---
[bingode](https://github.com/karashiiro/bingode)|![Code Size](https://img.shields.io/github/languages/code-size/karashiiro/bingode?label=%20)|A binary-packed data provider. Large and requires updates each patch, but works without relying on any websites besides The Lodestone.

## Usage
Refer to the [`examples/`](examples) folder for basic usage.

## Contributing
Make sure to checkout the submodules if you are changing CSS selector information.

### Dependencies
  * [`go-bindata`](https://github.com/go-bindata/go-bindata)

### Repacking
To repack the submodules, just run `generate.sh`.
