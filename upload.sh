#!/usr/bin/env sh -e

if [ "$#" -ne 3 ]; then
    echo "Illegal number of parameters, usage: ./upload.sh DEPLOYMENT_BUCKET VERSION_DIR FILE"
    echo "example: ./upload.sh flowlab-no-artifact-private 0.1.0 log-group-retention_0.1.0.zip"
    exit 1
fi

BUCKET=${1}
VERSION=${2}
FILE=${3}

aws s3 cp ${FILE} s3://${BUCKET}/log-group-retention/v${VERSION}/ --profile cloudformation@flowlab-development
