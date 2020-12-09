package user

import (
	"chatroom2/model"
	"strings"
)

func ParseMessageAction(inputMsg string) *model.MessageActionModel {
	action := ""
	msg := inputMsg

	if n := strings.Index(inputMsg, "]"); n != -1 {
		action = inputMsg[0:n]
		msg = inputMsg[n+1:]
	}

	return &model.MessageActionModel{
		InputContent: inputMsg,
		Action:       action,
		Message:      msg,
	}
}
