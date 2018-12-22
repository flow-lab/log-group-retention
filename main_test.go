package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/flow-lab/dlog"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const requestID = "1-581cf771-a006649127e371903a2de979"

func TestProcessEvent(t *testing.T) {
	os.Setenv("RETENTION_IN_DAYS", "60")
	cwl := &mockCloudWatchLogsClient{}

	result, err := ProcessEvent(cwl, dlog.NewRequestLogger(requestID, "test"))

	assert.Nil(t, err)
	assert.Len(t, result, 1)
}

func TestGetLogGroups(t *testing.T) {
	cwl := &mockCloudWatchLogsClient{}

	logGroups, err := GetLogGroups(cwl)

	check(t, err)
	assert.NotNil(t, logGroups)
	assert.Equal(t, 2, len(logGroups))
}

func TestPutRetentionPolicy(t *testing.T) {
	cwl := &mockCloudWatchLogsClient{}

	var logGroups []LogGroup
	test := "test"
	logGroup := LogGroup{
		LogGroupName: &test,
	}
	logGroups = append(logGroups, logGroup)

	result, err := PutRetentionPolicy(logGroups, cwl, dlog.NewRequestLogger(requestID, "test"))

	assert.Nil(t, err)
	assert.Len(t, result, 1)
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
		panic(err)
	}
}

// Define a mock struct to be used in your unit tests of myFunc.
type mockCloudWatchLogsClient struct {
	cloudwatchlogsiface.CloudWatchLogsAPI
}

func (m *mockCloudWatchLogsClient) DescribeLogGroups(input *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	var inputJSON = readFile("testdata/describeLogGroups-output.json")
	var describeLogGroupsOutput cloudwatchlogs.DescribeLogGroupsOutput
	err := json.Unmarshal(inputJSON, &describeLogGroupsOutput)
	if err != nil {
		panic(err)
	}
	return &describeLogGroupsOutput, nil
}

func (m *mockCloudWatchLogsClient) DescribeLogGroupsPages(input *cloudwatchlogs.DescribeLogGroupsInput, f func(*cloudwatchlogs.DescribeLogGroupsOutput, bool) bool) error {
	var inputJSON = readFile("testdata/describeLogGroups-output.json")
	var describeLogGroupsOutput cloudwatchlogs.DescribeLogGroupsOutput
	err := json.Unmarshal(inputJSON, &describeLogGroupsOutput)
	if err != nil {
		panic(err)
	}
	f(&describeLogGroupsOutput, true)
	return nil
}

func (m *mockCloudWatchLogsClient) PutRetentionPolicy(input *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
	return &cloudwatchlogs.PutRetentionPolicyOutput{}, nil
}

func readFile(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return f
}
