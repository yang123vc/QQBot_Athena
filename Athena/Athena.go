package main

import (
	"QQBot_Athena/Athena/apis"
	"QQBot_Athena/Athena/models"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

var message = "又是熟悉的天花板\n"

func main() {
	c := cron.New()
	// weibo refresh
	c.AddFunc("0 */10 * * * ?", models.SendWeibo)

	// abyss reminding
	c.AddFunc("0 0 20 ? * 0,3", models.TimerAbyss)

	// setu ranking
	//c.AddFunc("0 0 8 * * ?",models.TimerRefreshSTRanking)
	c.Start()

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// 接收消息
	router.POST("/api/ReceiveMahuaOutput", apis.ReceiveMahuaOutput)
	// 增删 friends 数据库

	router.Run(":65321")
}

func init() {
	data := models.Msg{"2325839514", 2, "547902826", "547902826"}
	models.GetGM(data)
	models.SendMsg(data, message)
}
