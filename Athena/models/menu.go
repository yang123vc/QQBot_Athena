package models

// 读取菜单
// txt文件保存便于修改
func GetMenu(data Msg) {
	str := getText("menu.txt")
	SendMsg(data, str)
}
