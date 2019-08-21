package models

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var master string = "569927585"
var masterlist []string
var Flag_shutup bool = false

func Shutup(data Msg, msg string) {
	var ifRespone bool = false

	// 遍历管理列表
	for _, v := range masterlist {
		if v == data.MsgAct {
			ifRespone = true
			break
		}
	}

	if ifRespone || data.MsgAct == master {
		if msg[:6] == "禁言" {
			if strings.Contains(msg, "IR:at") {
				pos1 := strings.Index(msg, "[IR:at=")
				pos2 := strings.Index(msg, "]")
				obj := msg[pos1+7 : pos2]
				time, err := strconv.Atoi(msg[pos2+2:])
				if err != nil {
					SendMsg(data, "格式错误")
					return
				}
				shutupOne(data, obj, time*60)
				SendMsg(data, "如您所愿，Master")
				Flag_shutup = true
				return
			} else {
				return
			}
		} else if msg[:6] == "解禁" {
			if strings.Contains(msg, "IR:at") {
				pos1 := strings.Index(msg, "[IR:at=")
				pos2 := strings.Index(msg, "]")
				obj := msg[pos1+7 : pos2]
				shutupOne(data, obj, 0)
				SendMsg(data, "如您所愿，Master")
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}

func shutupOne(data Msg, obj string, time int) {
	sendJson := make(map[string]interface{})
	sendJson["响应qq"] = data.QQ
	sendJson["群号"] = data.MsgFrom
	sendJson["对象QQ"] = obj
	sendJson["时间"] = time

	bytesData, _ := json.Marshal(sendJson)

	url := "http://47.100.182.193:36524/api/v1/CleverQQ/Api_ShutUP"
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

func RefreshMasters(data Msg) {
	// 获取群管理
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

	// 更新自定义
	path := "txt\\master.txt"
	menuFile, fileError := os.Open(path)
	if fileError != nil {
		SendMsg(data, "列表不存在")
		return
	}
	inputReader := bufio.NewReader(menuFile)
	for {
		inputString, inputError := inputReader.ReadString('\n')
		masterlist = append(masterlist, inputString)
		if inputError == io.EOF {
			return
			SendMsg(data, "更新成功")
		}
	}
}

func AddMaster(mem []string) {
	for _, v := range mem {
		masterlist = append(masterlist, v)
	}
	return
}
