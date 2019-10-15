package dateRange

import (
	"fmt"
	"time"
)

type LastYearData struct{}

func (y LastYearData) String() string {
	now := time.Now()
	return fmt.Sprintf("%d-01-01", now.Year()-1)
}
