#!/usr/bin/env bash
RUN_NAME="ntool"
export GO111MODULE=on
go mod vendor
go build -o output/bin/${RUN_NAME}