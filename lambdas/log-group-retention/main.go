package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/flow-lab/dlog"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// LogGroup dto, used in processing
type LogGroup struct {
	LogGroupName    *string
	RetentionInDays *int64
}

// Handler for lambda execution
func Handler(ctx context.Context, d events.CloudWatchEvent) (string, error) {
	lambdaContext, _ := lambdacontext.FromContext(ctx)
	requestLogger := dlog.NewRequestLogger(lambdaContext.AwsRequestID, "log-group-retention")

	sess := session.Must(session.NewSession())
	client := cloudwatchlogs.New(sess, &aws.Config{})

	_, err := ProcessEvent(client, requestLogger)
	if err != nil {
		requestLogger.Errorf("unable to complete: %v", err)
		panic(fmt.Errorf("unable to complete: %v", err))
	}

	return "event processed", nil
}

// ProcessEvent gets log groups and puts retention policy
func ProcessEvent(logs cloudwatchlogsiface.CloudWatchLogsAPI, log *log.Entry) ([]string, error) {
	logGroups, err := GetLogGroups(logs)
	if err != nil {
		return nil, fmt.Errorf("get log groups: %v", err)
	}

	mappedLogGroups, _ := mapToLogGroups(logGroups)
	result, err := PutRetentionPolicy(mappedLogGroups, logs, log)
	if err != nil {
		return nil, fmt.Errorf("get log with missing subscriptions: %v", err)
	}

	return result, nil
}

// GetLogGroups gets all logs groups in account
func GetLogGroups(logs cloudwatchlogsiface.CloudWatchLogsAPI) ([]*cloudwatchlogs.LogGroup, error) {
	var logGroups []*cloudwatchlogs.LogGroup
	input := cloudwatchlogs.DescribeLogGroupsInput{}
	err := logs.DescribeLogGroupsPages(&input, func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
		for _, logGroup := range page.LogGroups {
			logGroups = append(logGroups, logGroup)
		}
		return true
	})
	if err != nil {
		return nil, fmt.Errorf("describe log Groups: %v", err)
	}
	return logGroups, nil
}

// PutRetentionPolicy puts retention policy if missing
func PutRetentionPolicy(logGroups []LogGroup, logs cloudwatchlogsiface.CloudWatchLogsAPI, log *log.Entry) ([]string, error) {
	var result []string
	for _, logGroup := range logGroups {
		if logGroup.RetentionInDays == nil {
			retentionInDays, err := strconv.ParseInt(os.Getenv("RETENTION_IN_DAYS"), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unable to parse RETENTION_IN_DAYS %s, %v", os.Getenv("RETENTION_IN_DAYS"), err)
			}
			input := cloudwatchlogs.PutRetentionPolicyInput{
				LogGroupName:    logGroup.LogGroupName,
				RetentionInDays: &retentionInDays,
			}

			log.Printf("put retention policy %s", *logGroup.LogGroupName)
			_, err = logs.PutRetentionPolicy(&input)
			if err != nil {
				return nil, fmt.Errorf("putRetentionPolicy for %s: %v", *logGroup.LogGroupName, err)
			}
			result = append(result, *logGroup.LogGroupName)
			log.Printf("putRetentionPolicy for %s", *logGroup.LogGroupName)
		}
	}
	return result, nil
}

func mapToLogGroups(groups []*cloudwatchlogs.LogGroup) ([]LogGroup, error) {
	var result []LogGroup
	for _, element := range groups {
		logGroup := LogGroup{
			LogGroupName:    element.LogGroupName,
			RetentionInDays: element.RetentionInDays,
		}
		result = append(result, logGroup)
	}
	return result, nil
}

func main() {
	lambda.Start(Handler)
}
