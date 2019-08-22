package models

import (
	"math/rand"
	"strconv"
	"time"
)

var Flag_pixiv bool = true
var Flag_One bool = true

var groupRanking = make(map[string]*ranking)

type ranking struct {
	rank map[string]int
}

func Pixiv(data Msg) {
	//msg := "[IR:pic=http://laoliapi.cn/king/tupian/2cykj]"
	/*
		计数部分
		if _, ok := groupRanking[data.MsgFrom]; !ok {
			groupRanking[data.MsgFrom] = &ranking{}
		}
		if _, ok := groupRanking[data.MsgFrom].rank[data.MsgAct]; !ok {
			groupRanking[data.MsgFrom].rank[data.MsgAct] = 1
		}

		groupRanking[data.MsgFrom].rank[data.MsgAct] = groupRanking[data.MsgFrom].rank[data.MsgAct] + 1

	*/
	if Flag_pixiv == true {
		msg := "[IR:pic=https://api.pixivic.com/illust]"
		SendMsg(data, msg)
		timer := time.NewTimer(2 * time.Second)
		Flag_pixiv = false
		<-timer.C
		Flag_pixiv = true
	} else {
		msg := "太快啦，Athena受不了的"
		SendMsg(data, msg)
	}
	return
}

func OneSeTu(data Msg) {

	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch rand1.Intn(7) {
	case 0:
		msg := "[IR:pic=https://s0.xinger.ink/acgimg/acgurl.php]"
		SendMsg(data, "0"+msg)
		return
	case 1:
		msg := "[IR:pic=https://sotama.cool/picture]"
		SendMsg(data, "1"+msg)
		return
	case 2:
		msg := "[IR:pic=http://www.dmoe.cc/random.php]"
		SendMsg(data, "2"+msg)
		return
	case 3:
		msg := "[IR:pic=http://laoliapi.cn/king/tupian/2cykj]"
		SendMsg(data, "3"+msg)
		return
	case 4:
		msg := "[IR:pic=http://acg.bakayun.cn/randbg.php]"
		SendMsg(data, "4"+msg)
		return
	case 5:
		msg := "[IR:pic=https://acg.toubiec.cn/random]"
		SendMsg(data, "5"+msg)
		return
	case 6:
		msg := "[IR:pic=http://pic.tsmp4.net/api/erciyuan/img.php]"
		SendMsg(data, "6"+msg)
		return
	case 7:
		msg := "[IR:pic=http://www.dmoe.cc/random.php]"
		SendMsg(data, "7"+msg)
		return
	}
	return
}

func Bing(data Msg) {
	msg := "[IR:pic=http://laoliapi.cn/king/tupian/biying.php]"
	SendMsg(data, msg)
	return
}

func GetSTList(data Msg) {
	if _, ok := groupRanking[data.MsgFrom]; !ok {
		str := "今日本群还没有人要过瑟图哦"
		str += "[IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\setu.gif]"
		SendMsg(data, str)
		return
	}
	str := "今日本群共有" + strconv.Itoa(len(groupRanking[data.MsgFrom].rank)) + "人索要过瑟图，他们是：\n"
	for k, v := range groupRanking[data.MsgFrom].rank {
		str += "[IR:at=" + k + "] 共计" + strconv.Itoa(v) + "张\n"
	}
	str += "[IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\setu.gif]"
	SendMsg(data, str)
}

func GetSTRanking(data Msg) {
	if _, ok := groupRanking[data.MsgFrom]; !ok {
		str := "今日本群还没有人要过瑟图哦"
		str += "[IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\setu.gif]"
		SendMsg(data, str)
		return
	}
	str := "今日本群共有" + strconv.Itoa(len(groupRanking[data.MsgFrom].rank)) + "人索要过瑟图\n"

	type num struct {
		name string
		num  int
	}

	first := num{}
	second := num{}
	third := num{}

	for k, v := range groupRanking[data.MsgFrom].rank {
		if v > third.num {
			if v > second.num {
				if v > first.num {
					third.num = second.num
					third.name = second.name
					second.num = first.num
					second.name = first.name
					first.num = v
					first.name = k
				} else {
					third.num = second.num
					third.name = second.name
					second.num = v
					second.name = k
				}
			} else {
				third.num = v
				third.name = k
			}
		}
	}

	if first.name == "" {
		SendMsg(data, "今天还没有人要过瑟图哦")
		return
	}
	if second.name == "" {
		SendMsg(data, "今天只有一个人要过瑟图哦\n[IR:at="+first.name+"] 共计"+strconv.Itoa(first.num)+"张"+"[IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\setu.gif]")
		return
	}
	if third.name == "" {
		SendMsg(data, "今天只有两个人要过瑟图哦\n[IR:at="+first.name+"] 共计"+strconv.Itoa(first.num)+"张\n[IR:at="+second.name+"] 共计"+strconv.Itoa(second.num)+"张"+"[IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\setu.gif]")
		return
	}
	str += "其中前三名为：\n"
	str += "1.[IR:at=" + first.name + "] 共计" + strconv.Itoa(first.num) + "张"
	str += "2.[IR:at=" + second.name + "] 共计" + strconv.Itoa(second.num) + "张"
	str += "3.[IR:at=" + third.name + "] 共计" + strconv.Itoa(third.num) + "张"

	str += "[IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\setu.gif]"
	SendMsg(data, str)
}
