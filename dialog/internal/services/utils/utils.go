package utils

import "time"

func ParseTime(t []byte) (time.Time, error) {

	return time.Parse("2006-01-02 15:04:05.000000", string(t))
}
