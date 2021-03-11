# godestone

[![Documentation Badge](https://img.shields.io/badge/docs-pkg.go.dev-007D9C)](https://pkg.go.dev/github.com/xivapi/godestone/v2)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/28006e7fe175446db0fd8d38c92795b7)](https://app.codacy.com/gh/karashiiro/godestone?utm_source=github.com&utm_medium=referral&utm_content=karashiiro/godestone&utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/xivapi/godestone/v2)](https://goreportcard.com/report/github.com/xivapi/godestone/v2)
[![Dependencies](https://img.shields.io/librariesio/github/xivapi/godestone)](https://libraries.io/github/xivapi/godestone)
[![Repo Size](https://img.shields.io/github/repo-size/xivapi/godestone)](https://github.com/xivapi/godestone/v2)

Go scraper for The Lodestone.

## Installation
Just add the package to your `go.mod` or run `go get github.com/xivapi/godestone/v2`.

## Usage
Refer to the [`examples/`](examples) folder for basic usage.

This library does not come with a data backing library (outside of testing). An instance of a data backing service is required to initialize the scraper.
Currently, the only available data backing library is [bingode](https://github.com/karashiiro/bingode); more will be available in the future.

## Contributing
Make sure to checkout the submodules if you are changing CSS selector information.

### Dependencies
  * [`go-bindata`](https://github.com/go-bindata/go-bindata)

### Repacking
To repack the submodules, just run `generate.sh`.
