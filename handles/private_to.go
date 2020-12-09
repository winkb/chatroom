package handles

import (
	"chatroom2/model"
	"chatroom2/public"
	"strings"
)

type PrivateToHandle struct {
}

func (q *PrivateToHandle) IsAccept(model *model.MessageActionModel) bool {
	return strings.HasPrefix(model.InputContent, "@")
}

func parseUserNameAndContent(input string) (name, content string) {
	s := strings.TrimLeft(input, "@")
	index := strings.Index(s, " ")

	if index == -1 {
		return
	}

	name = s[0:index]
	content = s[index+1:]

	return
}

func (q *PrivateToHandle) Handle(m *model.MessageActionModel, user *model.UserClientModel) {
	var privateUser *model.UserClientModel

	name, content := parseUserNameAndContent(m.InputContent)

	public.GetUserManager().Range(func(user *model.UserClientModel) {
		if user.Name == name {
			privateUser = user
		}
	})

	if privateUser == nil {
		public.PushPublicMessage(user.Name + ":" + content)
		return
	}

	privateUser.Message <- user.Name + ":" + content
}
