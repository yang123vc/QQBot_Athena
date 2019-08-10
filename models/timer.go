package models

import (
	"fmt"
	"github.com/robfig/cron"
)

func Timer1() {
	c := cron.New()
	spec := "0 20 * * 3,7 ?"

	c.AddFunc(spec, timerHandler1)
	c.Start()
	select {}
}

func Timer2() {
	c := cron.New()
	spec := "30 22 * * 3,7 ?"

	c.AddFunc(spec, timerHandler2)
	c.Start()
	select {}
}

func timerHandler1() {
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
