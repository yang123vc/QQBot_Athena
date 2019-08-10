package models

import "time"

var Flag_pixiv bool = true
var Flag_One bool = true

func Pixiv(data Msg) {
	//msg := "[IR:pic=http://laoliapi.cn/king/tupian/2cykj]"
	if Flag_pixiv == true {
		msg := "[IR:pic=https://api.pixivic.com/illust]"
		SendMsg(data, msg)
		timer := time.NewTimer(2 * time.Second)
		Flag_pixiv = false
		<-timer.C
		Flag_pixiv = true
	} else {
		msg := "太快啦，Athena受不了的"
		SendMsg(data, msg)
	}
	return
}

func OneSeTu(data Msg) {
	/*
		fileName, err := GetWebFile("https://s0.xinger.ink/acgimg/acgurl.php")
		if err != nil {
			SendMsg(data, "file error")
			return
		}
		fs, err := ReadFile(fileName)
		if err != nil {
			SendMsg(data, "Read error")
			return
		}
		UploadPic(data, fs)

	*/
	msg := "[IR:pic=https://s0.xinger.ink/acgimg/acgurl.php]"
	SendMsg(data, msg)
	return
}

func Bing(data Msg) {
	msg := "[IR:pic=http://laoliapi.cn/king/tupian/biying.php]"
	SendMsg(data, msg)
	return
}
