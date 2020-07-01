package gotime

import (
	"fmt"
	"time"
)

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeBefore(time1,time2 string) bool {
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		return true
	}
	return false
}

func NowTimeint() string {
	  return fmt.Sprintf("%v", time.Now().Unix())
}
