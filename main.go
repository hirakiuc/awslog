package main

import (
	"os"

	_ "github.com/hirakiuc/awslog/internal/command/groups"
	"github.com/hirakiuc/awslog/internal/options"
)

func main() {
	if _, err := options.ParseOptions(); err != nil {
		os.Exit(1)
	}
}
