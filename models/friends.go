package models

import (
	"Athena/database"
	"fmt"
)

func InsertIOS(name, uid string) {
	db, err := database.ConnectDB("friends")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	_ = db.QueryRow("insert into ios values(\"" + uid + "\",\"" + name + "\")")
	return
}

func InsertAndroid(name, uid string) {
	db, err := database.ConnectDB("friends")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	_ = db.QueryRow("insert into android values(\"" + uid + "\",\"" + name + "\")")
	return
}

func GetIOS(data Msg) {
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

func GetAndroid(data Msg) {
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
