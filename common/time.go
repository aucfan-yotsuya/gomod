package common

import "time"

func ParseDuration(s string) time.Duration {
	var d time.Duration
	d, _ = time.ParseDuration(s)
	return d
}
func NowJST() time.Time {
	return time.Now().In(time.FixedZone("JST", 9*3600))
}
func NextMonthFirstDay() time.Time {
	var n = NowJST()
	return time.Date(
		n.Year(), n.Month(), 1,
		0, 0, 0, 0,
		time.FixedZone("JST", 9*3600),
	).AddDate(0, 1, 0)
}
