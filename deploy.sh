#!/usr/bin/env sh -e

if [[ "$#" -ne 2 ]]; then
    echo "Illegal number of parameters, usage: ./deploy.sh RETENTION_IN_DAYS"
    echo "example: ./deploy.sh 1"
    echo "RETENTION_IN_DAYS valid values are: [1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 3653]"
    exit 1
fi

RETENTION_IN_DAYS=${1}

aws cloudformation deploy \
   --stack-name log-group-retention \
   --template-file cloudformation/template.yml \
   --parameter-overrides RetentionInDays=${RETENTION_IN_DAYS} \
   --capabilities CAPABILITY_IAM \
   --profile cloudformation@flowlab-development
