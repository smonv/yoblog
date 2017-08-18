#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yoblog cmd/yoblog/main.go

docker build -t tthanh/yoblog .

rm yoblog
