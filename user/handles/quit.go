package handles

import (
	"chatroom2/user/model"
)

type QuitHandle struct {
}

func (q *QuitHandle) IsAccept(model *model.MessageActionModel) bool {
	return model.Action == "quit"
}

func (q *QuitHandle) Handle(model *model.MessageActionModel, user *model.UserClientModel) {
	close(user.Message)
	user.IsQuit <- true
}
