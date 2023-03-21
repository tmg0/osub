package parser

import "time"

func Interval(inter string) (time.Duration, error) {
	multiplier := 1
	if len(inter) > 1 {
		switch inter[len(inter)-1] {
		case 's':
			inter = inter[:len(inter)-1]
		case 'm':
			multiplier = 60
			inter = inter[:len(inter)-1]
		case 'h':
			multiplier = 60 * 60
			inter = inter[:len(inter)-1]
		case 'd':
			multiplier = 24 * 60 * 60
			inter = inter[:len(inter)-1]
		}
	}
	duration, err := time.ParseDuration(inter)
	if err != nil {
		return 0, err
	}
	return time.Duration(multiplier) * duration, nil
}
