package main

import (
	"Athena/apis"
	"Athena/models"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	// weibo refresh
	c.AddFunc("0 */10 * * * ?", models.SendWeibo)

	// abyss reminding
	c.AddFunc("0 0 20 ? * 0,3", models.TimerHandler1)
	//"0 0 20 ? * 0,3"
	c.Start()

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// 接收消息
	router.POST("/api/ReceiveMahuaOutput", apis.ReceiveMahuaOutput)
	// 增删 friends 数据库

	router.Run(":65321")
}
