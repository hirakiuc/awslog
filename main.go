package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	_ "github.com/hirakiuc/awslog/internal/options"
)

func main() {
	session := session.New(
		&aws.Config{
			Region: aws.String("ap-northeast-1"),
		},
	)

	service := cloudwatchlogs.New(session)

	params := &cloudwatchlogs.DescribeLogGroupsInput{
		Limit: aws.Int64(10),
	}

	res, err := service.DescribeLogGroups(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(res)
}
