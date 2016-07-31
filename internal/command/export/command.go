package export

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hirakiuc/awslog/internal/awslogs"
	"github.com/hirakiuc/awslog/internal/options"
	"github.com/hirakiuc/awslog/internal/parser"
)

type Command struct {
	StartTime string `short:"s" long:"start" default:"now" description:"start time to export log events."`
	EndTime   string `short:"e" long:"end"   default:"now" description:"end time to export log events."`

	Args struct {
		LogGroupName  string `positional-arg-name:"GroupName"  description:"target LogGroup Name"`
		LogStreamName string `positional-arg-name:"StreamName" description:"target LogStream Name"`
	} `positional-args:"yes" required:"yes"`
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

func (c *Command) requestParams() (awslogs.LogEventsParams, error) {
	params := awslogs.NewLogEventsParams()

	params.Limit = 1000

	timeParser := parser.NewTimeTextParser(time.Now())
	v, err := timeParser.Parse(c.StartTime)
	if err != nil {
		return params, err
	}
	params.StartTime = v

	v, err = timeParser.Parse(c.EndTime)
	if err != nil {
		return params, err
	}
	params.EndTime = v

	params.LogGroupName = c.Args.LogGroupName
	params.LogStreamName = c.Args.LogStreamName

	return params, nil
}

func (c *Command) Execute(args []string) error {
	service := awslogs.NewAwsLogs()

	params, err := c.requestParams()
	if err != nil {
		return err
	}

	err = service.LogEvents(params, func(logEvent *cloudwatchlogs.OutputLogEvent, lastEntry bool) bool {
		fmt.Println(*logEvent.Message)
		return !lastEntry
	})

	if err != nil {
		fmt.Println(err)
	}
	return err
}
