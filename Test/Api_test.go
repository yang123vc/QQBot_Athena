package Test

import (
	"QQBot_Athena/Athena/database"
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	var ifRespone bool = false
	var masterlist []string
	db, err := database.ConnectDB("group")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("SELECT * FROM `group`.`547902826` WHERE sponsor=?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}

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

		masterlist = append(masterlist, qq)
	}

	masterlist = append(masterlist, "456234")
	masterlist = append(masterlist, "412334")
	masterlist = append(masterlist, "451534")
	masterlist = append(masterlist, "4124")
	masterlist = append(masterlist, "123")
	// 遍历管理列表
	for _, v := range masterlist {
		if v == "102799448" {
			ifRespone = true
			break
		}
	}
	fmt.Println(ifRespone)
}
