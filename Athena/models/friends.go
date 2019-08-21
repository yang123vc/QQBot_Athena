package models

import (
	"QQBot_Athena/Athena/database"
	"fmt"
	"strings"
)

// 数据库添加内容
func insertIOS(name, uid string) {
	db, err := database.ConnectDB("friends")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	_ = db.QueryRow("insert into ios values(\"" + uid + "\",\"" + name + "\")")
	return
}
func insertAndroid(name, uid string) {
	db, err := database.ConnectDB("friends")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	_ = db.QueryRow("insert into android values(\"" + uid + "\",\"" + name + "\")")
	return
}

// 数据库获取内容
func getIOS(data Msg) {
	db, _ := database.ConnectDB("friends")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM ios")
	if err != nil {
		fmt.Print(err)
		SendMsg(data, "数据库错误")
		return
	}
	msg := ""
	for rows.Next() { //满足条件依次下一层
		var uid string
		var name string

		rows.Columns()

		err = rows.Scan(&uid, &name)
		msg += name + "  "
		msg += uid + "\n"
	}
	SendMsg(data, msg)
	return
}
func getAndroid(data Msg) {
	db, err := database.ConnectDB("friends")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM android")
	if err != nil {
		fmt.Print(err)
		SendMsg(data, "数据库错误")
		return
	}
	msg := ""
	var uid string
	var name string
	for rows.Next() { //满足条件依次下一层
		rows.Columns()

		err = rows.Scan(&uid, &name)
		fmt.Println(name)
		msg += name + "  "
		fmt.Println(uid)
		msg += uid + "\n"
	}
	SendMsg(data, msg)
	return
}

// 游戏好友相关操作
func IOS(data Msg, msg string) {
	if msg == "ios好友" {
		getIOS(data)
		return
	}
	if len(msg) > 10 {
		if msg[:9] == "添加ios" {
			if strings.Contains(msg[10:], "+") {
				str := strings.Split(msg[10:], "+")
				if len(str) != 2 {
					SendMsg(data, "格式错误")
					return
				}
				insertIOS(str[0], str[1])
				SendMsg(data, "添加成功")
				return
			}
		}
	}
}
func Android(data Msg, msg string) {
	if msg == "ios好友" {
		getAndroid(data)
		return
	}
	if len(msg) > 10 {
		if msg[:12] == "添加ios" {
			if strings.Contains(msg[13:], "+") {
				str := strings.Split(msg[13:], "+")
				if len(str) != 2 {
					SendMsg(data, "格式错误")
					return
				}
				insertAndroid(str[0], str[1])
				SendMsg(data, "添加成功")
				return
			}
		}
	}
}
