package main

import (
	"Athena/CleverQQ"
	"Athena/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestBing(t *testing.T) {
	data := models.Msg{"2325839514", 2, "892301661", "569927585"}
	models.Bing(data)
	return
}

func TestPixiv(t *testing.T) {
	data := models.Msg{"2325839514", 2, "892301661", "569927585"}
	msg := "[IR:pic=https://api.pixivic.com/illust]"
	msg = "[IR:pic=https://s0.xinger.ink/acgimg/acgurl.php]"
	models.SendMsg(data, msg)
}

func TestGetWebFile(t *testing.T) {
	models.GetWebFile("https://s0.xinger.ink/acgimg/acgurl.php")
}

func TestUploadPic(t *testing.T) {
	data := models.Msg{"2325839514", 2, "892301661", "569927585"}
	fs, _ := models.ReadFile("resource\\4597115562435754637.jpg")
	models.UploadPic(data, fs)
}

func TestGin(t *testing.T) {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		type Json struct {
			Data string `json:"data"`
		}

		var json Json

		fmt.Println(json.Data)
	})
}

func TestWeibo(t *testing.T) {
	//models.TestRefresh()
	models.SendWeibo()
	//models.WeiboTimer()
}

func TestTest(t *testing.T) {
	CleverQQ.CleverQQ_init("47.100.182.193:36524")
	data := CleverQQ.Data{ResQQ: "2325839514", RecQQ: "752390981"}
	CleverQQ.PBHomeWork(data, "test_name", "test_topic", "test_content")
}
