package util

import (
	"fmt"
	"time"
)

const readTimeFormat = "2006年1月2日"
const writeTimeFormat = "1/2"

var weekdayMap = []string{
	"日",
	"月",
	"火",
	"水",
	"木",
	"金",
	"土",
}

// Time means customized time.Time
type Time struct {
	time.Time
}

// JST is a timezone for "Asia/Tokyo"
var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

// MoesiaParseInLocation is customized version of time.ParseInLocation
func MoesiaParseInLocation(str string) (t *Time, err error) {
	date, err := time.ParseInLocation(readTimeFormat, str, JST)
	if err != nil {
		return
	}
	t = &Time{date}
	return
}

// MoesiaFormat is customized version of time.Format
func (t *Time) MoesiaFormat() (str string) {
	str = fmt.Sprintf("%s (%s)", t.Format(writeTimeFormat), weekdayMap[t.Weekday()])
	return
}
