# godestone-scraper

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/28006e7fe175446db0fd8d38c92795b7)](https://app.codacy.com/gh/karashiiro/godestone?utm_source=github.com&utm_medium=referral&utm_content=karashiiro/godestone&utm_campaign=Badge_Grade)
[![Dependencies](https://img.shields.io/librariesio/github/karashiiro/godestone)](https://libraries.io/github/karashiiro/godestone)
[![Repo Size](https://img.shields.io/github/repo-size/karashiiro/godestone)](https://github.com/karashiiro/godestone)

Go scraper for The Lodestone.

## Development
To generate FlatBuffer bindings for the data exports submodule, [flatc](https://google.github.io/flatbuffers) is required. Run `gen-flatbuf-bindings.sh` to create the bindings.

To package the submodules as assets, [go-bindata](https://github.com/go-bindata/go-bindata) is required. With that, run `pack-selectors.sh` to build the selectors file, and `pack-dats.sh` to build the game data file.
