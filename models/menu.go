package models

func GetMenu(data Msg) {
	str := getText("menu.txt")
	SendMsg(data, str)
}
