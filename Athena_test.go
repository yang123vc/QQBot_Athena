package main

import (
	"Athena/models"
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
