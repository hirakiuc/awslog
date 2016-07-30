package streams

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hirakiuc/awslog/internal/awslogs"
	"github.com/hirakiuc/awslog/internal/options"
)

type Command struct {
	LogStreamNamePrefix string `short:"p" long:"prefix" description:"LogStream name prefix" default:""`

	Args struct {
		LogGroupName string `positional-arg-name:"GroupName" description:"target LogGroup Name"`
	} `positional-args:"yes" required:"yes"`
}

var command Command

func init() {
	command = Command{}

	_, err := options.AddCommand(
		"streams",
		"Show Log Streams",
		"Show Log Streams on CloudWatchLogs",
		&command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *Command) requestParams() awslogs.LogStreamsParams {
	params := awslogs.NewLogStreamsParams()

	params.LogGroupName = c.Args.LogGroupName
	if len(c.LogStreamNamePrefix) > 0 {
		params.LogStreamNamePrefix = c.LogStreamNamePrefix
	}

	return params
}

func (c *Command) validate(args []string) error {
	if len(c.Args.LogGroupName) == 0 {
		return errors.New("Require LogGroupName.")
	}

	return nil
}

func (c *Command) Execute(args []string) error {
	service := awslogs.NewAwsLogs()

	err := service.LogStreams(c.requestParams(), func(stream *cloudwatchlogs.LogStream, lastEntry bool) bool {
		fmt.Println(*stream.LogStreamName)
		return !lastEntry
	})

	if err != nil {
		fmt.Println(err)
	}
	return err
}
