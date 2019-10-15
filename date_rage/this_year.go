package dateRange

import (
	"fmt"
	"time"
)

type ThisYearData struct{}

func (y ThisYearData) String() string {
	now := time.Now()
	return fmt.Sprintf("%d-01-01", now.Year())
}
