package main

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"time"
)

func main() {
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.PRC,
		WeekStartsAt: carbon.Sunday,
		Locale:       "zh-CN",
	})

	// 将标准 time.Time 转换成 Carbon
	carbon.CreateFromStdTime(time.Now())
	// 将 Carbon 转换成标准 time.Time
	carbon.Now().StdTime()

	// 今天此刻
	fmt.Printf("%s", carbon.Now())

}
