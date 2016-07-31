package parser

import (
	"regexp"
	"strconv"
	"time"
)

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

const timeTextPattern = `\A\s?(\d+)\s?(m|minutes?|h|hours?|d|days?|w|weeks?)\s?(|ago)\z`
const nowPattern = `\A\s?(now)\s?\z`

var timeMap = map[string]int64{
	"m": 60,
	"h": 60 * 60,
	"d": 60 * 60 * 24,
	"w": 60 * 60 * 24 * 7,
}

type timeExpression struct {
	Num  int64
	Unit int64
}

func newTimeExpression(valueText string, unitText string) (*timeExpression, error) {
	val, err := strconv.ParseInt(valueText, 10, 64)
	if err != nil {
		return nil, err
	}

	return &timeExpression{
		Num:  val,
		Unit: timeMap[unitText[0:1]],
	}, nil
}

func (exp *timeExpression) timeFrom(t time.Time) time.Time {
	diff := time.Duration(-1*exp.Num*exp.Unit) * time.Second
	return t.Add(diff)
}

type TimeTextParser struct {
	now         time.Time
	timePattern *regexp.Regexp
	nowPattern  *regexp.Regexp
}

func NewTimeTextParser(now time.Time) TimeTextParser {
	return TimeTextParser{
		now:         now,
		timePattern: regexp.MustCompile(timeTextPattern),
		nowPattern:  regexp.MustCompile(nowPattern),
	}
}

func timeToMilliSeconds(t time.Time) int64 {
	return t.Unix() * 1000
}

func (parser *TimeTextParser) Parse(text string) (int64, error) {
	if isNow := parser.nowPattern.MatchString(text); isNow == true {
		return parser.nowTime()
	}

	result := parser.timePattern.FindStringSubmatch(text)
	if len(result) == 4 {
		return parser.diffTime(result)
	}

	return parser.formattedTime(text)
}

// nowTime return unixtime in Millisecond
func (parser *TimeTextParser) nowTime() (int64, error) {
	return timeToMilliSeconds(parser.now), nil
}

// diffTime create a time with diff by now and return unixtime in Millisecond
func (parser *TimeTextParser) diffTime(matches []string) (int64, error) {
	timeDiff, err := newTimeExpression(matches[1], matches[2])
	if err != nil {
		return 0, err
	}

	return timeToMilliSeconds(timeDiff.timeFrom(parser.now)), nil
}

// formattedTime parse text and return unixtime in Millisecond
func (parser *TimeTextParser) formattedTime(text string) (int64, error) {
	const FORMAT = "2006-01-02 15:04:05 -0700"
	t, err := time.Parse(FORMAT, text)
	if err != nil {
		return 0, err
	}

	return timeToMilliSeconds(t), nil
}
