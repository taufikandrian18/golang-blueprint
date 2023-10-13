package utility

import (
	"strconv"
)

func ParseStringToFloat(stringFloat, format string) (f float64) {
	f, _ = strconv.ParseFloat(stringFloat, 64)
	return
}

func ParseFloatToString(f float64) (stringFloat string) {
	stringFloat = strconv.FormatFloat(f, 'f', -1, 64)
	return
}

func IterateNumber(start, end int) (data []int) {
	for i := start; i <= end; i++ {
		data = append(data, i)
	}
	return
}
