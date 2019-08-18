package CleverQQ

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var Host string = "localhost:36524"

type Data struct {
	ResQQ string
	RecQQ string
	ObjQQ string
}

const (
	Msg_Friend   int = 1
	Msg_Group    int = 2
	Msg_Forum    int = 3
	Msg_TemGroup int = 4
	Msg_TemForum int = 5
)

func CleverQQ_init(IPwithPort string) {
	Host = IPwithPort
	return
}

func post(sendJson map[string]interface{}, apiName string) bool {
	bytesData, _ := json.Marshal(sendJson)

	url := "http://" + Host + "/api/v1/CleverQQ/" + apiName
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false
		fmt.Println("Post Failed")
	}
	return true
}

// ResQQ 发送qq
// RecQQ 接收qq
// voice 语音数据 []byte
func SendVoice(data Data, voice []byte) bool {
	sendJson := make(map[string]interface{})
	sendJson["响应QQ"] = data.ResQQ
	sendJson["接收QQ"] = data.RecQQ
	sendJson["语音数据"] = voice

	return post(sendJson, "Api_SendVoice")
}

//ResQQ 响应qq
//RecQQ 群号
//ObjQQ 对象qq
//set true设置管理 false取消管理
func SetAdmin(data Data, set bool) bool {
	sendJson := make(map[string]interface{})
	sendJson["响应QQ"] = data.ResQQ
	sendJson["群号"] = data.RecQQ
	sendJson["对象QQ"] = data.ObjQQ
	sendJson["操作方式"] = set

	return post(sendJson, "Api_SetAdmin")
}

//ResQQ 响应QQ
//RecQQ 群号
//name 作业名
//topic 标题
//content 内容
func PBHomeWork(data Data, name, topic, content string) bool {
	sendJson := make(map[string]interface{})
	sendJson["响应QQ"] = data.ResQQ
	sendJson["群号"] = data.RecQQ
	sendJson["作业名"] = name
	sendJson["标题"] = topic
	sendJson["内容"] = content

	return post(sendJson, "Api_PBHomeWork")
}

// ResQQ 响应qq
// RecQQ 群号
func GetGroupMenberNum(data Data) bool {
	sendJson := make(map[string]interface{})
	sendJson["响应QQ"] = data.ResQQ
	sendJson["群号"] = data.RecQQ

	return post(sendJson, "Api_GetGroupMemberNum")
}

//resQQ 响应qq
//msgType 信息类型 MSG_
//recQQ 群，讨论组，好友
//msg 内容
//boob 气泡
func SendMsg(data Data, msgType int, msg string, boob int) bool {
	sendJson := make(map[string]interface{})

	if msgType == 1 || msgType == 0 {
		sendJson["响应QQ"] = data.ResQQ
		sendJson["信息类型"] = msgType
		sendJson["收信对象群_讨论组"] = ""
		sendJson["收信QQ"] = data.RecQQ
		sendJson["内容"] = msg
		sendJson["气泡ID"] = boob
	} else {
		sendJson["响应QQ"] = data.ResQQ
		sendJson["信息类型"] = msgType
		sendJson["收信对象群_讨论组"] = data.RecQQ
		sendJson["收信QQ"] = data.RecQQ
		sendJson["内容"] = msg
		sendJson["气泡ID"] = boob
	}

	return post(sendJson, "Api_SendMsg")
}
