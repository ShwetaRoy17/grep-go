#!/bin/bash

set -e # stop on first error

mkdir -p bin

go build -o bin/grep-go ./app


./bin/grep-go "$@"