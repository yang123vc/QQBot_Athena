package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// MsgType
const (
	MSG_T_PERSON = 1
	MSG_T_GROUP  = 2
)

type Msg struct {
	QQ      string
	MsgType int
	MsgFrom string
	MsgAct  string
}

func SendMsg(data Msg, msg string) {
	sendJson := make(map[string]interface{})
	sendJson["响应qq"] = data.QQ
	sendJson["信息类型"] = data.MsgType
	sendJson["收信对象群_讨论组"] = data.MsgFrom
	sendJson["收信qq"] = data.MsgFrom
	sendJson["内容"] = msg
	sendJson["气泡ID"] = -1

	bytesData, _ := json.Marshal(sendJson)

	url := "http://47.100.182.193:36524/api/v1/CleverQQ/Api_SendMsg"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("send Failed")
	}
	return
}

func UploadPic(data Msg, pic []byte) {
	//encodeString := base64.StdEncoding.EncodeToString(pic)
	//fmt.Println(encodeString)

	sendJson := make(map[string]interface{})
	sendJson["响应qq"] = data.QQ
	sendJson["上传类型"] = data.MsgType
	sendJson["参考对象"] = data.MsgFrom
	sendJson["图片数据"] = pic

	bytesData, _ := json.Marshal(sendJson)

	url := "http://47.100.182.193:36524/api/v1/CleverQQ/Api_UpLoadPic"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		fmt.Println("send Failed")
	}
	return
}

func GetGroupAdmin(data Msg) {
	sendJson := make(map[string]interface{})
	sendJson["响应qq"] = data.QQ
	sendJson["群号"] = data.MsgFrom
	bytesData, _ := json.Marshal(sendJson)
	url := "http://47.100.182.193:36524/api/v1/CleverQQ/Api_GetGroupAdmin"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	return
}

func GetGM(data Msg) {
	if data.MsgFrom != "547902826" {
		SendMsg(data, "不支持本群")
	}
	sendJson := make(map[string]interface{})
	sendJson["响应qq"] = data.QQ
	sendJson["群号"] = data.MsgFrom
	bytesData, _ := json.Marshal(sendJson)
	url := "http://47.100.182.193:36524/api/v1/CleverQQ/Api_GetGroupMemberList"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	return
}

func SendJson(data Msg, content string) {
	sendJson := make(map[string]interface{})
	sendJson["响应qq"] = data.QQ
	sendJson["发送方式"] = 1
	sendJson["信息类型"] = data.MsgType
	sendJson["收信对象群_讨论组"] = data.MsgFrom
	sendJson["收信qq"] = data.MsgAct
	sendJson["json结构"] = content

	bytesData, _ := json.Marshal(sendJson)

	url := "http://47.100.182.193:36524/api/v1/CleverQQ/Api_SendJSON"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("send Failed")
	}
	return
}
