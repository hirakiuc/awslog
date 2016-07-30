package export

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hirakiuc/awslog/internal/awslogs"
	"github.com/hirakiuc/awslog/internal/options"
	"github.com/hirakiuc/awslog/internal/parser"
)

type Command struct {
	LogGroupName  string `short:"g" long:"group" description:"LogGroup name" required:"true"`
	LogStreamName string `short:"s" long:"stream" description:"LogStream name" required:"true"`
	From          string `short:"f" long:"from" description:"export log events from"`
	To            string `short:"t" long:"to" description:"export log events to"`
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

	if len(c.From) > 0 || len(c.To) > 0 {
		timeParser := parser.NewTimeTextParser(time.Now())
		timeParser.Parse(c.From)
		timeParser.Parse(c.To)
		return errors.New("finish")
	}

	err := service.LogEvents(c.requestParams(), func(logEvent *cloudwatchlogs.OutputLogEvent, lastEntry bool) bool {
		fmt.Println(*logEvent.Message)
		return !lastEntry
	})

	if err != nil {
		fmt.Println(err)
	}
	return err
}
