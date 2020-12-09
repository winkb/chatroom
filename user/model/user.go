package model

type UserClientModel struct {
	Name     string
	Message  chan string
	IsQuit   chan bool
	IsActive chan bool
}
