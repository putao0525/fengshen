package common

import (
	"fengshen/vars"
	"fmt"
	"github.com/6tail/lunar-go/calendar"
)

// JiaZi 生成60甲子
func JiaZi() []string {
	jiaZi := make([]string, 60)
	for i := 0; i < 60; i++ {
		jiaZi[i] = vars.TianGan[i%10] + vars.DiZhi[i%12]
	}
	return jiaZi
}

func ShiJian(year, month, day, hour int) string {
	return fmt.Sprintf("%d年%d月%d日%d时", year, month, day, hour)
}

func GanZhi(year, month, day, hour, minute int) *calendar.Lunar {
	l := calendar.NewLunar(year, month, day, hour, minute, 0)
	return l
}
