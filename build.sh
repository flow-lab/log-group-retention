#!/usr/bin/env sh -e

if [[ "$#" -ne 1 ]]; then
    echo "Illegal number of parameters, usage: ./build.sh version"
    echo "example: ./build.sh 0.1.0"
    exit 1
fi

VERSION=${1}

go get -d ./...
GOOS=linux go build -o main
ZIP_FILE=log-group-retention_${VERSION}.zip
zip ${ZIP_FILE} main
