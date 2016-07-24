package options

import (
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Debug bool `short:"D" long:"debug" description:"Show debug messages"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func init() {
	options = Options{}
}

func ParseOptions() ([]string, error) {
	return parser.Parse()
}

func GetOptions() *Options {
	return &options
}

func AddCommand(command string, shortDescription string, longDescription string, data interface{}) (*flags.Command, error) {
	return parser.AddCommand(command, shortDescription, longDescription, data)
}

func (opts Options) Validate() error {
	return nil
}
