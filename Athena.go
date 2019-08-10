package main

import (
	"Athena/apis"
	"Athena/models"
	"github.com/gin-gonic/gin"
)

func main() {

	go models.Timer1()
	go models.Timer2()
	// Debug
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// 接收消息
	router.POST("/api/ReceiveMahuaOutput", apis.ReceiveMahuaOutput)
	// 增删 friends 数据库

	router.Run(":65321")
}
