package date_range

import (
	"fmt"
	"time"
)

type FiveYearData struct{}

func (y FiveYearData) String() string {
	now := time.Now()
	return fmt.Sprintf("%d-01-01", now.Year()-5)
}
