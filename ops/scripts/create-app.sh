#!/bin/sh

set -e

cd src

GOOS=linux GOARCH=amd64 go build -o ../go31ngo

cd ..

cp go31ngo ../31ngo/ops/docker/go/go31ngo

rm -f go31ngo
