package models

import (
	"fmt"
	"github.com/robfig/cron"
)

func TimerHandler1() {
	//547902826
	data := Msg{"2325839514", 2, "547902826", "547902826"}
	SendMsg(data, "Athena提醒\n该打深渊啦")
	return
}

func timerHandler2() {
	data := Msg{"2325839514", 2, "547902826", "547902826"}
	SendMsg(data, "Athena提醒\n上线收水晶啦")
	return
}

func TimerTest() {
	c := cron.New()
	spec := "*/1 * * * * ?"

	c.AddFunc(spec, func() {
		fmt.Println("Test ok")
		return
	})

	c.Start()
	select {}
}
