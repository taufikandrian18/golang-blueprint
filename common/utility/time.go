package utility

import (
	"time"
)

func ParseStringToTime(stringDate, format string) (t time.Time) {
	t, _ = time.Parse(format, stringDate)
	return
}

func ParseTimeToString(t time.Time, format string) (stringDate string) {
	stringDate = t.Format(format)
	return
}
