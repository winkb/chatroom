package main

import (
	"chatroom2/user"
	"chatroom2/user/handles"
	"chatroom2/user/model"
	"chatroom2/user/public"
	"fmt"
	"net"
	"time"
)

var userManager *user.UserManager
var pubMessage chan string

func init() {
	userManager = user.NewUserManager()
	pubMessage = public.GetPublicMessage()

	public.SetUserManager(userManager)
}

func manager() {
	for {
		msg := <-pubMessage
		userManager.PublicMessage(msg)
	}
}

func handleClientConn(accept net.Conn) {
	defer accept.Close()

	addr := accept.RemoteAddr()

	newUserClient := &model.UserClientModel{
		Name:     addr.String(),
		Message:  make(chan string),
		IsQuit:   make(chan bool), // 这里如果不make的话会导致阻塞，select监听不到，消息塞不进去
		IsActive: make(chan bool),
	}

	userManager.Add(newUserClient)

	message := user.NewMessage(accept, newUserClient)

	message.RegisterHandle(&handles.RenameHandle{})
	message.RegisterHandle(&handles.QuitHandle{})
	message.RegisterHandle(&handles.PrivateToHandle{})
	message.RegisterHandle(&handles.BroadcastHandle{})

	// 自己消息自己去处理，别烦
	go message.WriteToUserClient()
	go message.ListenUserClientMsg()

	pubMessage <- newUserClient.Name + "上线了"

	for {
		select {
		case <-newUserClient.IsQuit:
			fmt.Println(fmt.Sprintf("[%s]下线", newUserClient.Name))
			userManager.Remove(newUserClient)
			return
		case <-newUserClient.IsActive:
		case <-time.After(time.Second * 300):
			fmt.Println(fmt.Sprintf("超过5分钟没有互动：系统将[%s]踢下线", newUserClient.Name))
			return
		}
	}
}

func main() {

	// 建立监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8083")

	if err != nil {
		panic("启动出错：" + err.Error())
	}

	defer listener.Close()

	go manager()

	// 收消息
	for {
		accept, err := listener.Accept()

		if err != nil {
			fmt.Println("数据接收错误:\n", err)
			continue
		}

		go handleClientConn(accept)
	}

}
