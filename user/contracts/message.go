package contracts

import (
	"chatroom2/user/model"
)

type IMessageHandle interface {
	Handle(model *model.MessageActionModel, user *model.UserClientModel)
	IsAccept(model *model.MessageActionModel) bool
}
