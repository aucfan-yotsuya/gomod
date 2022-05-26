package common

import "time"

func ParseDuration(s string) time.Duration {
	var d time.Duration
	d, _ = time.ParseDuration(s)
	return d
}
