package models

var parrotMap = make(map[string]*parrotGroup)

type parrotGroup struct {
	strTemp      [3]string
	flagParroted bool
}

// 是否复读
func IfParrot(msg Msg, str string) bool {

	if _, ok := parrotMap[msg.MsgFrom]; !ok {
		parrotMap[msg.MsgFrom] = &parrotGroup{flagParroted: false}
		return false
	}

	parrotMap[msg.MsgFrom].strTemp[2] = parrotMap[msg.MsgFrom].strTemp[1]
	parrotMap[msg.MsgFrom].strTemp[1] = parrotMap[msg.MsgFrom].strTemp[0]
	parrotMap[msg.MsgFrom].strTemp[0] = str

	if parrotMap[msg.MsgFrom].strTemp[1] == parrotMap[msg.MsgFrom].strTemp[0] {
		if parrotMap[msg.MsgFrom].strTemp[2] == parrotMap[msg.MsgFrom].strTemp[1] {
			parrot(msg)
			parrotMap[msg.MsgFrom].flagParroted = true
		}
		return true
	} else {
		parrotMap[msg.MsgFrom].flagParroted = false
		return false
	}
}

// 复读
func parrot(msg Msg) {
	if parrotMap[msg.MsgFrom].flagParroted {
		return
	}
	SendMsg(msg, parrotMap[msg.MsgFrom].strTemp[0])
	return
}
