package user

import (
	"chatroom2/model"
	"sync"
)

var userLock sync.RWMutex

type UserManager struct {
	users map[string]*model.UserClientModel
}

func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[string]*model.UserClientModel),
	}
}

func (u *UserManager) Add(userClient *model.UserClientModel) {
	userLock.Lock()
	defer userLock.Unlock()

	// 不用重复加入到users
	if _, ok := u.users[userClient.Name]; ok {
		return
	}

	u.users[userClient.Name] = userClient
}

func (u *UserManager) Remove(userClient *model.UserClientModel) {
	userLock.Lock()
	defer userLock.Unlock()

	for i, currentUser := range u.users {
		if currentUser.Name == userClient.Name {
			delete(u.users, i)
		}
	}
}

func (u *UserManager) Range(fun func(user *model.UserClientModel)) {
	userLock.RLock()
	defer userLock.RUnlock()

	for _, currentUser := range u.users {
		fun(currentUser)
	}
}

func (u *UserManager) PublicMessage(message string) {
	u.Range(func(user *model.UserClientModel) {
		user.Message <- message
	})
}
