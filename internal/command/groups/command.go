package groups

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/hirakiuc/awslog/internal/awslogs"
	"github.com/hirakiuc/awslog/internal/options"
)

// Command describe "groups" command.
type Command struct {
	LogGroupNamePrefix string `short:"p" long:"prefix" description:"LogGroup name prefix" default:""`
}

var command Command

func init() {
	command = Command{}

	_, err := options.AddCommand(
		"groups",
		"Show LogGroups",
		"Show LogGroups on CloudWatchLogs",
		&command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *Command) requestParams() awslogs.LogGroupsParams {
	params := awslogs.NewLogGroupsParams()

	params.Limit = 30

	if len(c.LogGroupNamePrefix) > 0 {
		params.LogGroupNamePrefix = c.LogGroupNamePrefix
	}

	return params
}

// Execute fetch and show each LogGroup names.
func (c *Command) Execute(args []string) error {
	service := awslogs.NewAwsLogs()

	err := service.LogGroups(c.requestParams(), func(group *cloudwatchlogs.LogGroup, lastEntry bool) bool {
		fmt.Println(*group.LogGroupName)
		return !lastEntry
	})

	if err != nil {
		fmt.Println(err)
	}
	return err
}
