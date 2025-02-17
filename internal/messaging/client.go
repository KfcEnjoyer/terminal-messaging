package messaging

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
	"terminal-messaging/internal/utils"
	"time"
)

type User interface {
	GetConnection() (*rpc.Client, error)
	Register(client *rpc.Client) error
}

type UserService struct {
	Username   string
	Address    string
	muteGlobal bool
	Client     *rpc.Client
}

func (us *UserService) GetConnection() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil
}

func (us *UserService) Register(username string) {
	var reply string
	if err := us.Client.Call("ServerService.Register", username, &reply); err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}

func (us *UserService) GetUsers() {
	var reply []string

	if err := us.Client.Call("ServerService.GetUsers", "", &reply); err != nil {
		fmt.Println(err)
	}
	for i := range reply {
		fmt.Println(reply[i])
	}
}

func MainMenu(us *UserService) {
	buffer := bufio.NewReader(os.Stdin)

	for {
		text, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		text = strings.TrimSpace(text)
		params := strings.Fields(text)

		switch params[0] {
		case "show":
			us.GetUsers()

		case "send":
			if err := utils.ValidateParams(params); err != nil {
				fmt.Println(err)
			} else {
				us.SendMessage(params)
			}

		case "global":
			if us.muteGlobal == false {
				us.SayGlobally(params)
			} else {
				fmt.Println("You muted global chat! You cannot send messages")
			}

		case "logout":
			us.LogOut()
			return

		case "create":
			us.CreateRoom(params)

		case "join":
			us.JoinRoom(params)

		case "leave":
			us.LeaveRoom(params)

		case "mute":
			us.MuteGlobalChat()

		case "unmute":
			us.UnmuteGlobalChat()

		case "showm":
			us.showMuted()

		case "sendroom":
			us.SendRoom(params)

		case "showroomus":
			us.ShowRoomUsers(params)

		//case "ban":
		//	us.Ban(params)

		case "stop":
			us.Stop()
			return
		}
	}
}

func (us *UserService) SendMessage(params []string) {
	var reply NoReply

	msg := Message{
		From:    us.Username,
		Target:  params[1],
		Content: strings.Join(params[2:], " "),
	}

	if err := us.Client.Call("ServerService.SendMessage", msg, &reply); err != nil {
		fmt.Println(err)
	}
}

func (us *UserService) SayGlobally(params []string) {
	var reply NoReply

	msg := new(Message)
	msg.From = us.Username
	msg.Content = strings.Join(params[1:], " ")

	if err := us.Client.Call("ServerService.SayGlobally", msg, &reply); err != nil {
		fmt.Println(err)
	}
}

func (us *UserService) SendRoom(params []string) {
	var reply NoReply

	msg := Message{
		From:    us.Username,
		Target:  params[1],
		Content: strings.Join(params[2:], " "),
	}

	if err := us.Client.Call("ServerService.SendRoom", msg, &reply); err != nil {
		fmt.Println(err)
	}
}

func (us *UserService) ReadMessages() {
	var reply []string

	for {
		if err := us.Client.Call("ServerService.ReadMessages", us.Username, &reply); err != nil {
			fmt.Println(err)
		}

		for i := range reply {
			fmt.Println(reply[i])
		}

		time.Sleep(time.Second)
	}
}

func (us *UserService) LogOut() {
	var reply string

	if err := us.Client.Call("ServerService.LogOut", us.Username, &reply); err != nil {
		fmt.Println(err)
	}

	if err := us.Client.Close(); err != nil {
		fmt.Println(err)
	}
}

func (us *UserService) Stop() {
	var reply NoReply

	if err := us.Client.Call("ServerService.StopServer", "", &reply); err != nil {
		fmt.Println(err)
	}

}

func (us *UserService) CreateRoom(params []string) {
	var reply string

	room := new(Room)
	room.Name = params[1]
	room.Owner = us.Username
	room.Users = append(room.Users, room.Owner)

	if err := us.Client.Call("ServerService.CreateRoom", room, &reply); err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}

func (us *UserService) ShowRoomUsers(params []string) {
	var reply string

	if err := us.Client.Call("ServerService.ShowRoomUsers", params, &reply); err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}

func (us *UserService) JoinRoom(params []string) {
	var reply string

	room := Room{
		Name:        params[1],
		UserJoining: us.Username,
	}

	if err := us.Client.Call("ServerService.JoinRoom", room, &reply); err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}

func (us *UserService) LeaveRoom(params []string) {
	var reply string

	params = append(params, us.Username)

	if err := us.Client.Call("ServerService.LeaveRoom", params, &reply); err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}

func (us *UserService) MuteGlobalChat() {
	var reply string

	if err := us.Client.Call("ServerService.MuteGlobal", us.Username, &reply); err != nil {
		fmt.Println(err)
	}

	us.muteGlobal = true

	fmt.Println(reply)
}

func (us *UserService) UnmuteGlobalChat() {
	var reply string

	if err := us.Client.Call("ServerService.UnmuteGlobal", us.Username, &reply); err != nil {
		fmt.Println(err)
	}

	us.muteGlobal = false

	fmt.Println(reply)
}

func (us *UserService) showMuted() {
	var reply string

	if err := us.Client.Call("ServerService.ShowMuted", "", &reply); err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}

//func (us *UserService) Ban(params []string) {
//	var reply string
//
//	params = append(params, us.Username)
//
//	if err := us.Client.Call("ServerService.Ban", params, &reply); err != nil {
//		fmt.Println(err)
//	}
//
//
//	fmt.Println(reply)
//}
