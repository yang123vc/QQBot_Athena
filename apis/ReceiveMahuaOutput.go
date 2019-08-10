package apis

import (
	"Athena/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type ReceiveJson struct {
	Result            string `json:"Result"`
	CreateTime        string `json:"CreateTime"`
	EventAdditionType int    `json:"EventAdditionType"`
	EventOperator     string `json:"EventOperator"`
	EventType         int    `json:"EventType"`
	FromNum           string `json:"FromNum"`
	JSON              string `json:"Json"`
	Message           string `json:"Message"`
	MessageID         string `json:"MessageId"`
	MessageNum        string `json:"MessageNum"`
	Platform          int    `json:"Platform"`
	RawMessage        string `json:"RawMessage"`
	ReceiverQq        string `json:"ReceiverQq"`
	Triggee           string `json:"Triggee"`
	TypeCode          string `json:"TypeCode"`
}

func ReceiveMahuaOutput(c *gin.Context) {
	var json ReceiveJson
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json error")
	}

	// 确定事件
	switch json.TypeCode {

	case "Api_UploadPicApiOut":
		// 上传图片
		pic := models.Msg{json.ReceiverQq, json.EventType, json.FromNum, json.EventOperator}
		guid := json.Result
		models.SendMsg(pic, guid)

	case "Event":
		// 接收消息
		/*
			// 显示消息内容
			switch json.EventType {
			case 1:
				fmt.Println("好友消息：" + json.FromNum + ":" + json.Message)
			case 2:
				fmt.Println("群" + json.FromNum + ":" + json.EventOperator + ":" + json.Message)
			}

		*/

		if json.EventType == 2 {
			// 创建 Msg 结构体
			msg := models.Msg{json.ReceiverQq, json.EventType, json.FromNum, json.EventOperator}
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
					} else if json.Message[:12] == "添加官服" {
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
