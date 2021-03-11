#!/bin/sh

echo Clearing ./internal/pack/...
rm -rf ./internal/pack

echo Packing CSS selectors...
go-bindata -o internal/pack/css/selectors.go -prefix "lodestone-css-selectors/" -ignore="(LICENSE|README.md|.git)" lodestone-css-selectors/...
sed -i "s/package main/package css/g" internal/pack/css/selectors.go

echo Done!