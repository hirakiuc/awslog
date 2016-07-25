package parser

import (
	"fmt"
	"regexp"
	"time"
)

type TimeTextParser struct {
	now     time.Time
	pattern *regexp.Regexp
}

const timeTextPattern = `\A\s?(\d+)\s?(m|minutes?|h|hours?|d|days?|w|weeks?)\s?\z`

// N minutes
// Nm, Nminute, Nminutes
//
// N hours
// Nh, Nhour, Nhours
//
// N days
// Nd, Nday, Ndays
//
// N weeks
// Nw, Nweek, Nweeks

func NewTimeTextParser(now time.Time) TimeTextParser {
	return TimeTextParser{
		now:     now,
		pattern: regexp.MustCompile(timeTextPattern),
	}
}

func (parser *TimeTextParser) Parse(text string) (time.Time, error) {
	result := parser.pattern.FindStringSubmatch(text)
	if len(result) == 3 {
		fmt.Println(result[1])
		fmt.Println(result[2])
	}
	return parser.now, nil
}
