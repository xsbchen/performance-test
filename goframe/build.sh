#! /bin/bash

cd src

go mod download

GOAMD64=v3 go build -ldflags="-s -w" -o ../bin/app .