package awslogs

import (
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
	startTime := now.AddDate(0, -1, 0).Unix() * 1000
	endTime := now.Unix() * 1000

	return LogEventsParams{
		EndTime:       endTime,
		StartTime:     startTime,
		Limit:         0,
		LogGroupName:  "",
		LogStreamName: "",
		StartFromHead: true,
	}
}

func (params LogEventsParams) convert() *cloudwatchlogs.GetLogEventsInput {
	ret := &cloudwatchlogs.GetLogEventsInput{
		StartTime:     aws.Int64(params.StartTime),
		EndTime:       aws.Int64(params.EndTime),
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
			if len(output.Events) == 0 {
				return false
			}

			for i, logEvent := range output.Events {
				lastEvent := (lastPage && len(output.Events) == (i-1))
				ret := fn(logEvent, lastEvent)
				if ret == false {
					return false
				}
			}

			// Wait to avoid exhausting api call.
			time.Sleep(100 * time.Millisecond)

			return true
		})
}
