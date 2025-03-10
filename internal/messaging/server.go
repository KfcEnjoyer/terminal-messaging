package messaging

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"terminal-messaging/internal/utils"
)

type NoReply bool

type Server interface {
	StartServer(port string) error
	Register(username string, reply *string) error
	GetUsers(args string, reply *[]string) error
	SendMessage(msg Message, reply *NoReply) error
	SayGlobally(msg Message, reply *NoReply) error
	LogOut(username string, reply *string) error
	ReadMessages(username string, reply *[]string) error
	StopServer(args string, reply *NoReply) error
}

type ServerService struct {
	Port        string
	Users       []string
	MessageQ    map[string][]string
	Stop        chan bool
	MutedGlobal []string
	Rooms       []Room
}

func StartServer(s *ServerService) {
	if err := rpc.Register(s); err != nil {
		fmt.Println(err)
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", s.Port)
	if err != nil {
		fmt.Println(err)
	}

	go http.Serve(listen, nil)
}

func (s *ServerService) Register(username string, reply *string) error {
	if err := utils.ValidateUsername(username, s.Users, 3); err != nil {
		*reply = fmt.Sprintf(err.Error())
		return err
	}

	*reply = fmt.Sprintf("Welcome to the server %s \n", username)

	s.Users = append(s.Users, username)

	for i := range s.MessageQ {
		s.MessageQ[i] = append(s.MessageQ[i], fmt.Sprintf("User: %s joined the chat!", username))
	}

	s.MessageQ[username] = nil

	fmt.Println("User " + username + " joined the chat")
	fmt.Println(s.Users)

	return nil
}

func (s *ServerService) GetUsers(args string, reply *[]string) error {
	*reply = append(*reply, "List of active users: \n")

	for i, user := range s.Users {
		*reply = append(*reply, fmt.Sprintf("User %d: %s", i, user))
	}

	return nil
}

func (s *ServerService) SendMessage(msg Message, reply *NoReply) error {
	target := msg.Target
	from := msg.From
	content := msg.Content

	message := fmt.Sprintf("%s sent you: %s", from, content)

	s.MessageQ[target] = append(s.MessageQ[target], message)

	return nil
}

func (s *ServerService) SayGlobally(msg Message, reply *NoReply) error {
	content := msg.Content

	message := fmt.Sprintf("%s said to global chat: %s", msg.From, content)

	for i, _ := range s.Users {
		if s.Users[i] == msg.From {
			continue
		} else if isUserMuted(s.Users[i], s.MutedGlobal) {
			continue
		} else {
			s.MessageQ[s.Users[i]] = append(s.MessageQ[s.Users[i]], message)
		}
	}

	return nil
}

func isUserMuted(user string, mutedList []string) bool {
	for _, muted := range mutedList {
		if user == muted {
			return true
		}
	}
	return false
}

func (s *ServerService) LogOut(username string, reply *string) error {
	delete(s.MessageQ, username)

	for i := 0; i < len(s.Users); i++ {
		if s.Users[i] == username {
			s.Users = append(s.Users[:i], s.Users[i+1:]...)
			i--
		}
	}

	*reply = fmt.Sprintf("User %s has logged out", username)

	for i := range s.MessageQ {
		s.MessageQ[i] = append(s.MessageQ[i], *reply)
	}

	fmt.Printf("User %s has logged out\n", username)
	return nil
}

func (s *ServerService) ReadMessages(username string, reply *[]string) error {
	*reply = s.MessageQ[username]
	s.MessageQ[username] = nil
	return nil
}

func (s *ServerService) StopServer(args string, reply *NoReply) error {
	s.Stop <- false

	return nil
}

func (s *ServerService) CreateRoom(room Room, reply *string) error {
	s.Rooms = append(s.Rooms, room)

	*reply = fmt.Sprintf("user: %s created a room %s", room.Owner, room.Name)
	return nil
}

func (s *ServerService) JoinRoom(room Room, reply *string) error {
	roomName := room.Name
	user := room.UserJoining

	for i := range s.Rooms {
		if s.Rooms[i].Name == roomName {
			s.Rooms[i].Users = append(s.Rooms[i].Users, user)
			*reply = fmt.Sprintf("User %s added to room %s", user, roomName)
			return nil
		}
	}

	*reply = fmt.Sprintf("Room %s not found", roomName)
	return nil
}

func (s *ServerService) LeaveRoom(params []string, reply *string) error {
	roomName := params[1]
	user := params[2]

	roomIndex := -1
	for i := range s.Rooms {
		if s.Rooms[i].Name == roomName {
			roomIndex = i
			break
		}
	}

	if roomIndex == -1 {
		*reply = fmt.Sprintf("Room '%s' does not exist", roomName)
		return fmt.Errorf("room '%s' not found", roomName)
	}

	roomRef := &s.Rooms[roomIndex]

	userIndex := -1
	for i, roomUser := range roomRef.Users {
		if roomUser == user {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		*reply = fmt.Sprintf("User '%s' is not in room '%s'", user, roomName)
		return fmt.Errorf("user '%s' not in room '%s'", user, roomName)
	}

	roomRef.Users = append(roomRef.Users[:userIndex], roomRef.Users[userIndex+1:]...)

	for _, remainingUser := range roomRef.Users {
		if s.MessageQ[remainingUser] != nil {
			s.MessageQ[remainingUser] = append(s.MessageQ[remainingUser], fmt.Sprintf("User '%s' left room '%s'", user, roomName))
		}
	}

	return nil
}

func (s *ServerService) ShowRoomUsers(params []string, reply *string) error {
	roomName := params[1]

	room := Room{}

	for i := range s.Rooms {
		if s.Rooms[i].Name == roomName {
			room = s.Rooms[i]
		}
	}

	for i := range room.Users {
		*reply += "USer " + string(i) + ":" + " " + room.Users[i] + "\n"
	}

	return nil
}

func (s *ServerService) SendRoom(msg Message, reply *NoReply) error {
	var room Room

	for i := range s.Rooms {
		if s.Rooms[i].Name == msg.Target {
			room = s.Rooms[i]

		} else {

			return nil
		}
	}

	message := fmt.Sprintf("%s sent to room %s: %s", msg.From, msg.Target, msg.Content)

	for i := range room.Users {
		s.MessageQ[room.Users[i]] = append(s.MessageQ[room.Users[i]], message)
	}

	return nil
}

func (s *ServerService) DeleteRoom(roomMame, username string) error {

	var room Room

	for i := range s.Rooms {
		if s.Rooms[i].Name == roomMame {
			room = s.Rooms[i]
		}
	}

	if username == room.Owner {
		fmt.Println("Lmao ur the owner")
	} else {
		fmt.Println("Niigea")
	}

	return nil

}

func (s *ServerService) MuteGlobal(username string, reply *string) error {

	for i := range s.MutedGlobal {
		if s.MutedGlobal[i] == username {
			*reply = "You already muted global chat!"
			return nil
		}
	}

	s.MutedGlobal = append(s.MutedGlobal, username)

	*reply = "You successfully muted global chat"

	return nil
}

func (s *ServerService) UnmuteGlobal(username string, reply *string) error {

	for user := range s.MutedGlobal {
		if s.MutedGlobal[user] == username {
			s.MutedGlobal = append(s.MutedGlobal[:user], s.MutedGlobal[user+1:]...)
		}
	}

	*reply = "You successfully unmuted global chat"

	return nil
}

func (s *ServerService) ShowMuted(args string, reply *string) error {
	*reply = "Here is the list of muted users: \n"

	for i := range s.MutedGlobal {
		*reply += "user: " + s.MutedGlobal[i] + "\n"
	}

	return nil
}

//func (s *ServerService) Ban(params []string, reply *string) error {
//	sender := params[2]
//	username := params[1]
//
//	fmt.Println(sender)
//	fmt.Println(username)
//
//	if sender != "admin" {
//		*reply = "You are not admin"
//		return nil
//	}
//
//	for user := range s.Users {
//		if s.Users[user] == username {
//			s.Users = append(s.Users[:user], s.Users[user+1:]...)
//		}
//	}
//
//	*reply = "You banned: " + username
//
//	return nil
//}
