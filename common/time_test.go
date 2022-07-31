package common

import (
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextMonth(t *testing.T) {
	ymd := os.Getenv("NEXT_YMD")
	re := regexp.MustCompile(`(\d{4})\-(\d{2})\-(\d{2})`)
	yyyy := re.ReplaceAllString(ymd, `$1`)
	mm := re.ReplaceAllString(ymd, `$2`)
	dd := re.ReplaceAllString(ymd, `$3`)
	assert.Equal(t, NextMonth(), time.Date(
		Atoi(yyyy),
		time.Month(Atoi(mm)),
		Atoi(dd),
		0,
		0,
		0,
		0,
		time.FixedZone("JST", 9*3600),
	))
}
func TestPrevMonth(t *testing.T) {
	ymd := os.Getenv("PREV_YMD")
	re := regexp.MustCompile(`(\d{4})\-(\d{2})\-(\d{2})`)
	yyyy := re.ReplaceAllString(ymd, `$1`)
	mm := re.ReplaceAllString(ymd, `$2`)
	dd := re.ReplaceAllString(ymd, `$3`)
	assert.Equal(t, PrevMonth(), time.Date(
		Atoi(yyyy),
		time.Month(Atoi(mm)),
		Atoi(dd),
		0,
		0,
		0,
		0,
		time.FixedZone("JST", 9*3600),
	))
}
func TestCalcYearMonth(t *testing.T) {
	{
		n := time.Date(
			2022, 03, 30, 0, 0, 0, 0, time.FixedZone("JST", 9*3600),
		)
		assert.Equal(t, CalcYearMonth(n, 1), "202203")
		assert.Equal(t, CalcYearMonth(n, 29), "202203")
	}
	{
		n := time.Date(
			2020, 02, 29, 0, 0, 0, 0, time.FixedZone("JST", 9*3600),
		)
		assert.Equal(t, CalcYearMonth(n, 1), "202002")
		assert.Equal(t, CalcYearMonth(n, 28), "202002")
		assert.Equal(t, CalcYearMonth(n, 29), "202002")
		assert.Equal(t, CalcYearMonth(n, 30), "202001")
		assert.Equal(t, CalcYearMonth(n, 31), "202001")
	}
	{
		n := time.Date(
			2020, 01, 30, 0, 0, 0, 0, time.FixedZone("JST", 9*3600),
		)
		assert.Equal(t, CalcYearMonth(n, 29), "202001")
		assert.Equal(t, CalcYearMonth(n, 30), "202001")
		assert.Equal(t, CalcYearMonth(n, 31), "201912")
	}
}
