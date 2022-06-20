package common

import "time"

func ParseDuration(s string) time.Duration {
	var d time.Duration
	d, _ = time.ParseDuration(s)
	return d
}
func NowJST() time.Time {
	var jst *time.Location
	jst, _ = time.LoadLocation("Asia/Tokyo")
	return time.Now().In(jst)
}
