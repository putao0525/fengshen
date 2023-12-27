package common

import (
	"fmt"
	"testing"
)

func TestGanZhi(t *testing.T) {
	l := GanZhi(2023, 12, 26, 10, 30)
	fmt.Printf("%v  %v  %v  %v", l.GetYearInGanZhi(), l.GetMonthInGanZhi(), l.GetDayInGanZhi(), l.GetHour())
}
