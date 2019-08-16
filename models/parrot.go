package models

var StrTemp [3]string
var FlagParroted bool = false

func IfParrot(msg Msg, str string) {
	StrTemp[2] = StrTemp[1]
	StrTemp[1] = StrTemp[0]
	StrTemp[0] = str
	if StrTemp[1] == StrTemp[0] {
		if StrTemp[2] == StrTemp[1] {
			parrot(msg)
			FlagParroted = true
		}
	} else {
		FlagParroted = false
	}
}

func parrot(msg Msg) {
	if FlagParroted {
		return
	}
	SendMsg(msg, StrTemp[0])
	return
}
