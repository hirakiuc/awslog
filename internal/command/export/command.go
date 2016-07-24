package export

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hirakiuc/awslog/internal/awslogs"
	"github.com/hirakiuc/awslog/internal/options"
)

type Command struct {
	LogGroupName  string `short:"g" long:"group" description:"LogGroup name" required:"true"`
	LogStreamName string `short:"s" long:"stream" description:"LogStream name" required:"true"`
}

var command Command

func init() {
	command = Command{}

	_, err := options.AddCommand(
		"export",
		"Export LogEvents",
		"Export LogEvents in LogGroup and LogStream",
		&command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *Command) requestParams() awslogs.LogEventsParams {
	params := awslogs.NewLogEventsParams()

	params.Limit = 1000

	params.LogGroupName = c.LogGroupName
	params.LogStreamName = c.LogStreamName

	return params
}

func (c *Command) Execute(args []string) error {
	service := awslogs.NewAwsLogs()

	err := service.LogEvents(c.requestParams(), func(logEvent *cloudwatchlogs.OutputLogEvent, lastEntry bool) bool {
		fmt.Println(*logEvent.Message)
		return !lastEntry
	})

	if err != nil {
		fmt.Println(err)
	}
	return err
}
