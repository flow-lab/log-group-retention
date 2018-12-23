## AWS log group retention [![Build Status](https://travis-ci.org/flow-lab/log-group-retention.svg?branch=master)](https://travis-ci.org/flow-lab/log-group-retention) [![codecov](https://codecov.io/gh/flow-lab/log-group-retention/branch/master/graph/badge.svg)](https://codecov.io/gh/flow-lab/log-group-retention)

CloudWatch logs can be expensive if no retention is set for the log groups. This lambda function will fire up every 60 
minutes and iterate over all log groups in the AWS account and if RetentionPolicy is not present it will create policy 
with 60 days retention*.

* this value can be changed. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 
and 3653.

```


                         ------------------------------------------------
   CloudWatch Event     |   log-group-retention                          |
   (every 60 minutes)   |   1. Get all log groups                        |
----------------------> |   2. If retention policy is missing            |
                        |   3. Put retention policy                      |
                         ------------------------------------------------
```

Run`build.sh`, `upload.sh` and `deploy.sh` accordingly to get app up and running.

## CodePipeline

To deploy
```sh
aws cloudformation deploy \
    --stack-name log-group-retention-codepipeline \
    --parameter-overrides ApplicationName="log-group-retention" GitHubUser="flow-lab" GitHubRepository="log-group-retention" GitHubOAuthToken="GITHUB_TOKEN" \
    --role-arn "ROLE_ARN"
    --template cloudformation/pipeline.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --profile cloudformation@flowlab-development
```

## License

MIT License (MIT)