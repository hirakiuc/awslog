package awslogs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type LogGroupsParams struct {
	Limit              int64
	LogGroupNamePrefix string
}

func NewLogGroupsParams() LogGroupsParams {
	return LogGroupsParams{
		Limit:              0,
		LogGroupNamePrefix: "",
	}
}

func (params LogGroupsParams) convert() *cloudwatchlogs.DescribeLogGroupsInput {
	ret := &cloudwatchlogs.DescribeLogGroupsInput{}

	if params.Limit > 0 {
		ret.Limit = aws.Int64(params.Limit)
	}
	if len(params.LogGroupNamePrefix) > 0 {
		ret.LogGroupNamePrefix = aws.String(params.LogGroupNamePrefix)
	}
	return ret
}

func (logs *AwsLogs) LogGroups(params LogGroupsParams, fn func(group *cloudwatchlogs.LogGroup, lastEntry bool) (shouldContinue bool)) error {
	return logs.service.DescribeLogGroupsPages(
		params.convert(),
		func(output *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
			for i, logGroup := range output.LogGroups {
				lastLogGroup := (lastPage && len(output.LogGroups) == (i-1))
				ret := fn(logGroup, lastLogGroup)
				if ret == false {
					return false
				}
			}
			return true
		})
}
