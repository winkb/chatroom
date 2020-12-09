package user

import (
	. "chatroom2/user/contracts"
	"chatroom2/user/model"
	"fmt"
	"net"
)

type Message struct {
	conn    net.Conn
	user    *model.UserClientModel
	handles []IMessageHandle
}

func NewMessage(conn net.Conn, user *model.UserClientModel) *Message {
	return &Message{
		conn: conn,
		user: user,
	}
}

func (c *Message) RegisterHandle(h IMessageHandle) {
	c.handles = append(c.handles, h)
}

func (m *Message) ListenUserClientMsg() {
	clientBufData := make([]byte, 2049)

	for {
		n, err := m.conn.Read(clientBufData)

		if n == 0 {
			fmt.Println(m.user.Name, "已经下线")
			m.user.IsQuit <- true
			return
		}

		if err != nil {
			fmt.Println("数据读取错误:\n", err)
			continue
		}
		// 去掉最后的换行符
		msg := string(clientBufData[:n-1])

		messageActionModel := ParseMessageAction(msg)

		m.user.IsActive <- true

		for _, handle := range m.handles {
			if handle.IsAccept(messageActionModel) {
				handle.Handle(messageActionModel, m.user)
				break
			}
		}

	}
}

func (m *Message) WriteToUserClient() {
	for {
		toUserMsg, ok := <-m.user.Message
		if !ok {
			fmt.Println("message 消息通道关闭")
			return
		}

		_, err := m.conn.Write([]byte(toUserMsg + "\n"))

		if err != nil {
			fmt.Println(fmt.Sprintf("消息发送给用户[%s]失败:%s", m.user.Name, err))
		}
	}
}
