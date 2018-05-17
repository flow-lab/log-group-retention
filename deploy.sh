#!/usr/bin/env bash -e

if [ "$#" -ne 3 ]; then
    echo "Illegal number of parameters, usage: ./deploy.sh DEPLOYMENT_BUCKET DIST_FILE RETENTION_IN_DAYS"
    echo "example: ./deploy.sh deployment-bucket-s3bucket-rkjl6q60hsw8 deployment-201805171544.zip 15"
    echo "RETENTION_IN_DAYS valid values are: [1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 3653]"
    exit 1
fi

BUCKET=${1}
FILE=${2}
RETENTION_IN_DAYS=${3}

aws s3 cp ${FILE} s3://${BUCKET}/log-group-retention/ --profile cloudformation@flowlabdev

aws cloudformation deploy \
   --stack-name log-group-retention \
   --template-file cloudformation/template.yml \
   --parameter-overrides DeploymentBucket=${BUCKET} DeploymentFile=${FILE} RetentionInDays=${RETENTION_IN_DAYS} \
   --capabilities CAPABILITY_IAM \
   --profile cloudformation@flowlabdev
