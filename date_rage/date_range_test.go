package date_range_test

import (
	"fmt"
	"testing"
	"time"

	dateRange "github.com/jamieabc/finance_indicator/date_rage"
	"github.com/stretchr/testify/assert"
)

func TestNewRangeWhenToday(t *testing.T) {
	r := dateRange.NewRange(dateRange.Today)

	now := time.Now()
	today := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())

	assert.Equal(t, today, r.String(), "wrong today")
}

func TestNewRangeWhenThisYear(t *testing.T) {
	y := dateRange.NewRange(dateRange.ThisYear)

	now := time.Now()
	thisYear := fmt.Sprintf("%d-01-01", now.Year())

	assert.Equal(t, thisYear, y.String(), "wrong this year")
}

func TestNewRangeWhenLastYear(t *testing.T) {
	y := dateRange.NewRange(dateRange.LastYear)

	now := time.Now()
	lastYear := fmt.Sprintf("%d-01-01", now.Year()-1)

	assert.Equal(t, lastYear, y.String(), "wrong last year")
}
