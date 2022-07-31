package common

import (
	"fmt"
	"time"
)

func ParseDuration(s string) time.Duration {
	var d time.Duration
	d, _ = time.ParseDuration(s)
	return d
}
func NowJST() time.Time {
	return time.Now().In(time.FixedZone("JST", 9*3600))
}
func NextMonth() time.Time {
	var n = NowJST()
	return time.Date(
		n.Year(), n.Month(), n.Day(),
		0, 0, 0, 0,
		time.FixedZone("JST", 9*3600),
	).AddDate(0, 1, 0)
}
func PrevMonth() time.Time {
	var n = NowJST()
	return time.Date(
		n.Year(), n.Month(), n.Day(),
		0, 0, 0, 0,
		time.FixedZone("JST", 9*3600),
	).AddDate(0, -1, 0)
}
func NextMonthFirstDay() time.Time {
	var n = NowJST()
	return time.Date(
		n.Year(), n.Month(), 1,
		0, 0, 0, 0,
		time.FixedZone("JST", 9*3600),
	).AddDate(0, 1, 0)
}
func PrevMonthFirstDay() time.Time {
	var n = NowJST()
	return time.Date(n.Year(), n.Month(), 1, 0, 0, 0, 0, time.FixedZone("JST", 9*3600)).
		AddDate(0, -1, 0)
}
func CalcYearMonth(t time.Time, day int) string {
	var n time.Time
	if t.Day() < day {
		n = time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, time.FixedZone("JST", 9*3600))
		return fmt.Sprintf("%04d%02d", n.Year(), n.Month())
	}
	return fmt.Sprintf("%04d%02d", t.Year(), t.Month())
}
