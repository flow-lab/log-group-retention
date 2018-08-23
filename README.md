## AWS log group subscriber [![Build Status](https://travis-ci.org/flow-lab/log-group-retention.svg?branch=master)](https://travis-ci.org/flow-lab/log-group-retention) [![codecov](https://codecov.io/gh/flow-lab/log-group-retention/branch/master/graph/badge.svg)](https://codecov.io/gh/flow-lab/log-group-retention)

Lambda function that puts `RetentionPolicy` for log groups if missing.


```


                         ---------------------------------------
   CloudWatch Event     |   log-group-retention                 |
   (every 60 minutes)   |   1. Get all log groups               |
----------------------> |   2. If retention policy is missing   |
                        |   3. Put retention policy             |
                         ---------------------------------------
```


To build run commend below. It compiles sources to `main` binary file and zips
it to deployment package `deployment-123456789.zip`
```sh
./build.sh
```

To deploy to AWS with cloudformation template use `deploy.sh` script