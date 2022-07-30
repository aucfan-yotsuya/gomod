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
