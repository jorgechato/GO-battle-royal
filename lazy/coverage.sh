#!/usr/bin/env bash

set -e

for D in $(go list ./... | grep -v vendor); do
    go test -coverprofile=coverage.out $D
done
