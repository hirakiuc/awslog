package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type AwsLogs struct {
	service *cloudwatchlogs.CloudWatchLogs
}

func newSession() *session.Session {
	return session.New()
}

func newService() *cloudwatchlogs.CloudWatchLogs {
	return cloudwatchlogs.New(newSession())
}

func NewAwsLogs() *AwsLogs {
	return &AwsLogs{
		service: newService(),
	}
}
