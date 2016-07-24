package awslogs

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type LogEventsParams struct {
	EndTime       int64
	StartTime     int64
	Limit         int64
	LogGroupName  string
	LogStreamName string
	StartFromHead bool
}

func NewLogEventsParams() LogEventsParams {
	now := time.Now()
	startTime := now.Unix() * 1000
	endTime := now.AddDate(0, 0, -2).Unix() * 1000

	return LogEventsParams{
		Limit:         1000,
		EndTime:       startTime,
		StartTime:     endTime,
		LogGroupName:  "",
		LogStreamName: "",
		StartFromHead: false,
	}
}

func (params LogEventsParams) convert() *cloudwatchlogs.GetLogEventsInput {
	ret := &cloudwatchlogs.GetLogEventsInput{
		StartTime:     aws.Int64(params.StartTime),
		EndTime:       aws.Int64(params.EndTime),
		Limit:         aws.Int64(params.Limit),
		LogGroupName:  aws.String(params.LogGroupName),
		LogStreamName: aws.String(params.LogStreamName),
		StartFromHead: aws.Bool(params.StartFromHead),
	}

	return ret
}

func (logs *AwsLogs) LogEvents(params LogEventsParams, fn func(logEvent *cloudwatchlogs.OutputLogEvent, lastEntry bool) (shouldContinue bool)) error {
	return logs.service.GetLogEventsPages(
		params.convert(),
		func(output *cloudwatchlogs.GetLogEventsOutput, lastPage bool) bool {
			fmt.Println("last?", lastPage, ", size:", len(output.Events))
			if lastPage && len(output.Events) == 0 {
				return false
			}

			for i, logEvent := range output.Events {
				lastEvent := (lastPage && len(output.Events) == (i-1))
				ret := fn(logEvent, lastEvent)
				if ret == false {
					return false
				}
			}
			return true
		})
}
