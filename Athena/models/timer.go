package models

func TimerAbyss() {
	//547902826
	data := Msg{"2325839514", 2, "547902826", "547902826"}
	SendMsg(data, "Athena提醒\n该打深渊啦\n[IR:at=全体人员][IR:pic=C:\\Users\\Administrator\\Desktop\\client\\resource\\abyss.jpg]")
	return
}

func TimerRefreshSTRanking() {
	data := Msg{"2325839514", 2, "547902826", "569927585"}
	GetSTRanking(data)

	for k, _ := range groupRanking {
		delete(groupRanking, k)
	}

	SendMsg(data, "统计已清空")
	return
}
