package handles

import (
	"chatroom2/user/model"
	"chatroom2/user/public"
)

type RenameHandle struct {
}

func (q *RenameHandle) IsAccept(model *model.MessageActionModel) bool {
	return model.Action == "rename"
}

func (q *RenameHandle) Handle(msgModel *model.MessageActionModel, user *model.UserClientModel) {
	userManager := public.GetUserManager()
	isUsed := false
	newName := msgModel.Message

	userManager.Range(func(user *model.UserClientModel) {
		if user.Name == newName {
			isUsed = true
		}
	})

	if isUsed {
		user.Message <- "改名失败，用户名重复:" + user.Name
		return
	}

	user.Name = msgModel.Message
	user.Message <- "改名成功:" + user.Name
}
