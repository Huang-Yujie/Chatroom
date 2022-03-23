package chat

import (
	"log"

	"github.com/Huang-Yujie/Chatroom/global"
)

type broadcaster struct {
	users map[uint64]*User

	enteringChannel chan *User
	leavingChannel  chan *User
	messageChannel  chan *Message
}

var Broadcaster = &broadcaster{
	users: make(map[uint64]*User),

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChannel:  make(chan *Message, 1024),
}

func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enteringChannel:
			b.users[user.UserID] = user
		case user := <-b.leavingChannel:
			delete(b.users, user.UserID)
			user.CloseMessageChannel()
		case msg := <-b.messageChannel:
			for _, user := range b.users {
				if user.UserID == msg.User.UserID {
					continue
				}
				user.Send(msg)
			}
		}
	}
}

func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}

func (b *broadcaster) Broadcast(msg *Message) {
	if len(b.messageChannel) >= global.ChatroomSettings.MessageQueueLength {
		log.Println("broadcast queue 满了")
	}
	b.messageChannel <- msg
}
