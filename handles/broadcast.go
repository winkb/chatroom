package handles

import (
	"chatroom2/model"
	"chatroom2/public"
)

type BroadcastHandle struct {
}

func (q *BroadcastHandle) IsAccept(model *model.MessageActionModel) bool {
	return model.Action == ""
}

func (q *BroadcastHandle) Handle(model *model.MessageActionModel, user *model.UserClientModel) {
	public.PushPublicMessage(user.Name + ":" + model.Message)
}
