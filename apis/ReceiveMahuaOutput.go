package apis

import (
	"Athena/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
)

type ReceiveJson struct {
	Result            string `json:"Result"`
	CreateTime        string `json:"CreateTime"`
	EventAdditionType int    `json:"EventAdditionType"`
	// 事件发起者
	EventOperator string `json:"EventOperator"`
	// 事件类型
	EventType int `json:"EventType"`
	// 事件来源
	FromNum string `json:"FromNum"`
	JSON    string `json:"Json"`
	// 事件内容
	Message    string `json:"Message"`
	MessageID  string `json:"MessageId"`
	MessageNum string `json:"MessageNum"`
	Platform   int    `json:"Platform"`
	RawMessage string `json:"RawMessage"`
	// 响应QQ
	ReceiverQq string `json:"ReceiverQq"`
	// 事件响应者
	Triggee  string `json:"Triggee"`
	TypeCode string `json:"TypeCode"`
}

func ReceiveMahuaOutput(c *gin.Context) {
	var json ReceiveJson
	if err := c.ShouldBindJSON(&json); err != nil {
		// 接收格式错误
		//fmt.Println("json error")
		return
	}
	// 确定事件
	switch json.TypeCode {
	case "Event":
		msg := models.Msg{json.ReceiverQq, json.EventType, json.FromNum, json.EventOperator}

		switch json.EventType {
		case 203:
			msg.MsgType = 2
			switch rand.Intn(3) {
			case 0:
				models.SendMsg(msg, "虾仁猪心")
				return
			case 1:
				models.SendMsg(msg, "口球带好")
				return
			case 2:
				models.SendMsg(msg, "[Face178.gif][Face67.gif]")
				return
			case 3:
				models.SendMsg(msg, "发言不规范，群员两行泪")
			}
		case 2:
			switch json.Message {
			case "pixiv":
				models.Pixiv(msg)
			case "一张瑟图":
				models.OneSeTu(msg)
			case "/help":
				// /help
				models.GetMenu(msg)
			case "官服好友":
				// 显示官服好友:
				models.GetAndroid(msg)
			case "ios好友":
				// 显示ios好友
				models.GetIOS(msg)
			case "[IR:at=2325839514] ":
				models.SendMsg(msg, "/help 查询技能\nhttps://github.com/Logiase/QQBot_Athena.git")
			case "[IR:at=2325839514]":
				models.SendMsg(msg, "/help 查询技能\nhttps://github.com/Logiase/QQBot_Athena.git")
			default:
				// 添加内容
				if len(json.Message) > 13 {
					if json.Message[:9] == "添加ios" {
						if strings.Contains(json.Message[10:], "+") {
							str := strings.Split(json.Message[10:], "+")
							if len(str) != 2 {
								models.SendMsg(msg, "格式错误")
							}
							models.InsertIOS(str[0], str[1])
						}
					}
					if json.Message[:12] == "添加官服" {
						if strings.Contains(json.Message[13:], "+") {

							str := strings.Split(json.Message[13:], "+")
							if len(str) != 2 {
								models.SendMsg(msg, "格式错误")
							}
							models.InsertAndroid(str[0], str[1])
						}
					}
				}
			}
		}
	}
	return
}
