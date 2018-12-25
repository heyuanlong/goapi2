package service

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"

	kinit "goapi2/initialize"
	kcode "goapi2/work/code"
)

const SETTLE_TIME_SECOND = 1

func InitSystem() {

	t, err := kinit.Conf.GetInt("server", "is_test")
	if err == nil && t == 1 {
		kinit.LogInfo.Println("this is test server")
		kcode.IS_TEST_SERVER = t
	}

	fmt.Println("version:", 201810171632)
}

func Run() {
	t, err := kinit.Conf.GetInt("server", "timestamp")
	if err != nil {
		kinit.LogError.Println("get timestamp fail:", err)
		return
	}
	timestamp := int64(t)
	timer := GetTimer(timestamp)

	t1 := time.NewTimer(time.Second * SETTLE_TIME_SECOND)
	t3 := time.NewTimer(time.Second * time.Duration(timer))

	BusiSettle.Init()
	for {
		select {

		case <-t1.C:
			t1.Reset(time.Second * SETTLE_TIME_SECOND)
			BusiSettle.Run()

		case <-t3.C:
			timer := GetTimer(timestamp)
			t3.Reset(time.Second * time.Duration(timer))
			kinit.LogError.Println("深夜调用")
		}
	}
}

func GetTimer(timestamp int64) int64 {
	now.WeekStartDay = time.Monday // Set Monday as first day, default is Sunday
	nowTime := time.Now().Unix()
	beginDay := now.BeginningOfDay().Unix()
	var timer int64
	if (nowTime - beginDay) < timestamp {
		timer = timestamp - (nowTime - beginDay)
	} else {
		timer = (beginDay + 3600*24) - nowTime + timestamp
	}
	return timer
}
