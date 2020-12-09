package public

import (
	"chatroom2/user"
)

var pubMessage chan string
var userManager *user.UserManager

func init() {
	pubMessage = make(chan string)
}

func PushPublicMessage(message string) {
	pubMessage <- message
}

func GetPublicMessage() chan string {
	return pubMessage
}

func SetUserManager(mg *user.UserManager) {
	userManager = mg
}

func GetUserManager() *user.UserManager {
	return userManager
}
