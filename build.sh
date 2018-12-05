#!/usr/bin/env bash -e

go get -d ./...
GOOS=linux go build -o main
ZIP_FILE=deployment-"$(date +"%Y%m%d%H%M")".zip
zip ${ZIP_FILE} main
