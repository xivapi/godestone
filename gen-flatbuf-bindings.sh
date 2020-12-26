#!/bin/sh
flatc --go -o ./pack/exports --go-namespace pack --gen-onefile --filename-suffix "" ./lodestone-data-exports/schema/*.fbs