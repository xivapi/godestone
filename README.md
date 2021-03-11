# godestone

[![Documentation Badge](https://img.shields.io/badge/docs-pkg.go.dev-007D9C)](https://pkg.go.dev/github.com/xivapi/godestone/v2)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/28006e7fe175446db0fd8d38c92795b7)](https://app.codacy.com/gh/karashiiro/godestone?utm_source=github.com&utm_medium=referral&utm_content=karashiiro/godestone&utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/xivapi/godestone/v2)](https://goreportcard.com/report/github.com/xivapi/godestone/v2)
[![Dependencies](https://img.shields.io/librariesio/github/xivapi/godestone)](https://libraries.io/github/xivapi/godestone)
[![Code Size](https://img.shields.io/github/languages/code-size/xivapi/godestone)](https://github.com/xivapi/godestone)

Go scraper for The Lodestone.

## Installation
Just add the package to your `go.mod` or run `go get github.com/xivapi/godestone/v2`.

Also be sure to install a data provider service to initialize the scraper.

## Data providers
Package|Description
---|---
[bingode](https://github.com/karashiiro/bingode)|A binary-packed data provider. Large and requires updates each patch, but works without relying on any websites besides The Lodestone.

## Usage
Refer to the [`examples/`](examples) folder for basic usage.

## Contributing
Make sure to checkout the submodules if you are changing CSS selector information.

### Dependencies
  * [`go-bindata`](https://github.com/go-bindata/go-bindata)

### Repacking
To repack the submodules, just run `generate.sh`.
