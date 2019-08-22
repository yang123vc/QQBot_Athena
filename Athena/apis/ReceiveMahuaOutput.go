package apis

import (
	"QQBot_Athena/Athena/database"
	"QQBot_Athena/Athena/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
		return
	}
	// 确定事件
	switch json.TypeCode {
	case "Api_GetGroupMemberListApiOut":
		getGroupMember(json.Result)
		data := models.Msg{"2325839514", 2, "547902826", "569927585"}
		rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := rand1.Float64()/4 + 0.9
		models.SendMsg(data, "世界线变动率"+strconv.FormatFloat(num, 'f', 7, 64))
		return
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
			models.IfParrot(msg, json.Message)

			// 命令判断
			switch json.Message {
			case "一张瑟图":
				models.OneSeTu(msg)
			case "/help":
				models.GetMenu(msg)
			case "[IR:at=2325839514] ":
				models.SendMsg(msg, "/help 查询技能\nhttps://github.com/Logiase/QQBot_Athena.git")
			case "[IR:at=2325839514]":
				models.SendMsg(msg, "/help 查询技能\nhttps://github.com/Logiase/QQBot_Athena.git")
			case "世界线收束":
				if json.EventOperator == "569927585" {
					models.GetGM(msg)
					return
				}
				return
			case "刷新管理":
				models.RefreshMasters(msg)
				return
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

var db, _ = database.ConnectDB("group")

func getGroupMember(result string) {
	type GM struct {
		Ec        int    `json:"ec"`
		Errcode   int    `json:"errcode"`
		Em        string `json:"em"`
		AdmNum    int    `json:"adm_num"`
		AdmMax    int    `json:"adm_max"`
		Vecsize   int    `json:"vecsize"`
		Levelname struct {
			Num1   string `json:"1"`
			Num2   string `json:"2"`
			Num3   string `json:"3"`
			Num4   string `json:"4"`
			Num5   string `json:"5"`
			Num6   string `json:"6"`
			Num10  string `json:"10"`
			Num11  string `json:"11"`
			Num12  string `json:"12"`
			Num13  string `json:"13"`
			Num14  string `json:"14"`
			Num15  string `json:"15"`
			Num101 string `json:"101"`
			Num102 string `json:"102"`
			Num103 string `json:"103"`
			Num104 string `json:"104"`
			Num105 string `json:"105"`
			Num106 string `json:"106"`
			Num107 string `json:"107"`
			Num108 string `json:"108"`
			Num109 string `json:"109"`
			Num110 string `json:"110"`
			Num111 string `json:"111"`
			Num112 string `json:"112"`
			Num113 string `json:"113"`
			Num114 string `json:"114"`
			Num115 string `json:"115"`
			Num116 string `json:"116"`
			Num117 string `json:"117"`
			Num118 string `json:"118"`
			Num197 string `json:"197"`
			Num198 string `json:"198"`
			Num199 string `json:"199"`
		} `json:"levelname"`
		Mems []struct {
			Uin           int `json:"uin"`
			Role          int `json:"role"`
			Flag          int `json:"flag"`
			G             int `json:"g"`
			JoinTime      int `json:"join_time"`
			LastSpeakTime int `json:"last_speak_time"`
			Lv            struct {
				Point int `json:"point"`
				Level int `json:"level"`
			} `json:"lv"`
			Nick string `json:"nick"`
			Card string `json:"card"`
			Qage int    `json:"qage"`
			Tags string `json:"tags"`
			Rm   int    `json:"rm"`
		} `json:"mems"`
		Count       int `json:"count"`
		SvrTime     int `json:"svr_time"`
		MaxCount    int `json:"max_count"`
		SearchCount int `json:"search_count"`
	}
	var gm GM
	json.Unmarshal([]byte(result), &gm)

	// 遍历更新数据库
	for _, k := range gm.Mems {
		rows, err := db.Query("SELECT * FROM `group`.`547902826` WHERE QQ=?", k.Uin)
		if err != nil {
			fmt.Println(err)
		}
		ifEx := false
		for rows.Next() {
			ifEx = true
		}
		if ifEx {
			_, err = db.Exec("UPDATE `group`.`547902826` SET NickName=?,role=? WHERE QQ=?", k.Card, k.Role, k.Uin)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			_, err = db.Exec("INSERT INTO `group`.`547902826`(QQ,NickName,role) VALUES (?,?,?)", k.Uin, k.Card, k.Role)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}
