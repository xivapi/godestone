#!/bin/sh
flatc --go -o ./pack/exports --go-namespace exports --gen-onefile --filename-suffix "" ./lodestone-data-exports/schema/*.fbs