package models

import (
	"QQBot_Athena/Athena/database"
	"strings"
)

// 保持长连接
var db, err = database.ConnectDB("group")

func insertGame(QQ, name, uid string, server int) (err error) {
	rows, _ := db.Query("SELECT * FROM `group`.`547902826` WHERE QQ=?", QQ)
	ifExist := false
	for rows.Next() {
		ifExist = true
	}

	if ifExist {
		_, err := db.Query("UPDATE `group`.`547902826` SET UID=?,GameName=?,GameServer=? WHERE QQ=?", uid, name, server, QQ)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec("INSERT `group`.`547902826`(QQ, GameServer, GameName, UID) VALUES (?,?,?,?)", QQ, server, name, uid)
		if err != nil {
			return err
		}
	}
	return nil
}
func getGame(data Msg, server int) {
	rows, err := db.Query("SELECT * FROM `group`.`547902826` WHERE GameServer=?", server)
	if err != nil {
		SendMsg(data, "数据库读取错误")
		return
	}
	msg := ""

	for rows.Next() {
		qq := ""
		nn := ""
		sp := 0
		gs := 0
		gn := ""
		uid := ""
		ro := 0
		money := 0

		err = rows.Scan(&qq, &nn, &sp, &gs, &gn, &uid, &ro, &money)

		msg += gn + " " + uid + "\n"
	}

	msg = strings.TrimRight(msg, "\n")
	SendMsg(data, msg)
	return
}

func IOS(data Msg, msg string) {
	if msg == "ios好友" {
		getGame(data, 2)
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
				if insertGame(data.MsgAct, str[0], str[1], 2) == nil {
					SendMsg(data, "添加成功")
					return
				} else {
					SendMsg(data, "insert failed")
					return
				}
			}
		}
	}
}
func Android(data Msg, msg string) {
	if msg == "官服好友" {
		getGame(data, 1)
		return
	}
	if len(msg) > 10 {
		if msg[:12] == "添加官服" {
			if strings.Contains(msg[13:], "+") {
				str := strings.Split(msg[13:], "+")
				if len(str) != 2 {
					SendMsg(data, "格式错误")
					return
				}
				if insertGame(data.MsgAct, str[0], str[1], 1) == nil {
					SendMsg(data, "添加成功")
					return
				} else {
					SendMsg(data, "insert failed")
					return
				}
			}
		}
	}
}
