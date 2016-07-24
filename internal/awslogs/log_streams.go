package awslogs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type LogStreamsParams struct {
	Descending          bool
	Limit               int64
	LogGroupName        string
	LogStreamNamePrefix string
	OrderBy             string
}

func NewLogStreamsParams() LogStreamsParams {
	return LogStreamsParams{
		Descending:          false,
		Limit:               0,
		LogGroupName:        "",
		LogStreamNamePrefix: "",
		OrderBy:             "LogStreamName",
	}
}

func (params LogStreamsParams) convert() *cloudwatchlogs.DescribeLogStreamsInput {
	ret := &cloudwatchlogs.DescribeLogStreamsInput{}

	ret.Descending = aws.Bool(params.Descending)

	if params.Limit > 0 {
		ret.Limit = aws.Int64(params.Limit)
	}
	if len(params.LogGroupName) > 0 {
		ret.LogGroupName = aws.String(params.LogGroupName)
	}
	if len(params.LogStreamNamePrefix) > 0 {
		ret.LogStreamNamePrefix = aws.String(params.LogStreamNamePrefix)
	}
	ret.OrderBy = aws.String(params.OrderBy)

	return ret
}

func (logs *AwsLogs) LogStreams(params LogStreamsParams, fn func(stream *cloudwatchlogs.LogStream, lastEntry bool) (shouldContinue bool)) error {
	return logs.service.DescribeLogStreamsPages(
		params.convert(),
		func(output *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
			for i, logStream := range output.LogStreams {
				lastStream := (lastPage && len(output.LogStreams) == (i-1))
				ret := fn(logStream, lastStream)
				if ret == false {
					return false
				}
			}
			return true
		})
}
