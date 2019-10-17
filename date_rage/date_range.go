package date_range

import (
	"fmt"
)

// Ranger - with String method
type Ranger interface {
	fmt.Stringer
}

type rangeType int

const (
	Today = iota
	ThisYear
	LastYear
	FiveYear
)

// NewRange - return object with String method
func NewRange(r rangeType) Ranger {
	switch r {
	case Today:
		return TodayData{}
	case ThisYear:
		return ThisYearData{}
	case LastYear:
		return LastYearData{}
	case FiveYear:
		return FiveYearData{}
	}
	return ThisYearData{}
}
