package date_range

import (
	"fmt"
	"time"
)

type TodayData struct{}

func (t TodayData) String() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
}
