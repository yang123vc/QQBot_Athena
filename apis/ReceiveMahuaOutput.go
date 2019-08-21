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
	case "Api_GetGroupAdminApiOut":
		models.AddMaster(getGroupAdmin(json.Result))
	case "Event":
		msg := models.Msg{json.ReceiverQq, json.EventType, json.FromNum, json.EventOperator}

		switch json.EventType {
		case 203:
			msg.MsgType = 2
			// 结构待优化
			if models.Flag_shutup == false {
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
			} else {
				models.Flag_shutup = false
			}
		case 2:
			// 复读跳过命令判断
			if models.IfParrot(msg, json.Message) {
				return
			}

			// 命令判断
			switch json.Message {
			case "pixiv":
				models.Pixiv(msg)
			case "一张瑟图":
				models.OneSeTu(msg)
			case "/help":
				// /help
				models.GetMenu(msg)
				/*
					case "官服好友":
						models.GetAndroid(msg)
					case "ios好友":
						models.GetIOS(msg)
				*/
			case "[IR:at=2325839514] ":
				models.SendMsg(msg, "/help 查询技能\nhttps://github.com/Logiase/QQBot_Athena.git")
			case "[IR:at=2325839514]":
				models.SendMsg(msg, "/help 查询技能\nhttps://github.com/Logiase/QQBot_Athena.git")
			case "刷新管理":
				models.RefreshMasters(msg)
				return
				/*
					case "绅士列表":
						models.GetSTList(msg)
						return
					case "绅士排名":
						models.GetSTRanking(msg)
						return

				*/
			default:
				// ios相关操作
				if strings.Contains(json.Message, "ios") {
					models.IOS(msg, json.Message)
					return
				}
				// 官服相关操作
				if strings.Contains(json.Message, "官服") {
					models.Android(msg, json.Message)
					return
				}
				// 禁言相关
				if strings.Contains(json.Message, "禁") {
					models.Shutup(msg, json.Message)
					return
				}
			}
		}
	}
	return
}

func getGroupAdmin(result string) []string {
	//str := strings.Join(strings.Split(result, "\\"), "")
	pos := strings.Index(result, "gAdmins")
	pos1 := strings.Index(result[pos:], "[")
	pos2 := strings.Index(result[pos:], "]")
	//fmt.Println(str[pos+pos1+1 : pos+pos2])
	return strings.Split(result[pos+pos1+1:pos+pos2], ",")
}
