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

func TestStr(t *testing.T) {
	str := "{\"config\":{\"forward\":true,\"type\":\"normal\",\"autosize\":true},\"prompt\":\"[分享]假酒害人 !毛子沙雕时刻 不要笑挑战! 气氛突然苏维埃!\",\"app\":\"com.tencent.structmsg\",\"ver\":\"0.0.0.1\",\"view\":\"news\",\"meta\":{\"news\":{\"title\": \"假酒害人 !毛子沙雕时刻 不要笑挑战! 气氛突然苏维埃!政委..\",\"desc\":\"已观看16.2万次\",\"preview\":\"http:\\/\\/url.cn\\/5MxlHPR\",\"tag\":\"哔哩哔哩\",\"jumpUrl\":\"http:\\/\\/url.cn\\/5goSM5f\",\"appid\":100951776,\"app_type\":1,\"action\":\"\",\"source_url\":\"\",\"source_icon\":\"\",\"android_pkg_name\":\"\"}},\"desc\":\"新闻\"}"
	fmt.Println(str)
}
