package common

import "time"

func ParseDuration(s string) time.Duration {
	var d time.Duration
	d, _ = time.ParseDuration(s)
	return d
}
func NowJST() time.Time {
	return time.Now().In(time.FixedZone("JST", 9*60*60))
}
