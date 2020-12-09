package handles

import (
	"chatroom2/user/model"
	"chatroom2/user/public"
)

type BroadcastHandle struct {
}

func (q *BroadcastHandle) IsAccept(model *model.MessageActionModel) bool {
	return model.Action == ""
}

func (q *BroadcastHandle) Handle(model *model.MessageActionModel, user *model.UserClientModel) {
	public.PushPublicMessage(user.Name + ":" + model.Message)
}
