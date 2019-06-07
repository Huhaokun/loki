#/bin/bash

export GO111MODULE=on
go mod download
go build -o ./bin/app github.com/Huhaokun/loki/app
