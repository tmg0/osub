package resolve

import (
	"time"
)

func Interval(inter string) (time.Duration, error) {
	duration, err := time.ParseDuration(inter)
	if err != nil {
		return 0, err
	}
	return duration, nil
}
