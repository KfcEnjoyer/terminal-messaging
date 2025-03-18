package messaging_test

import (
	"terminal-messaging/internal/messaging"
	"testing"
	"time"
)

func TestServerServiceInitialization(t *testing.T) {
	s := new(messaging.ServerService)
	s.MessageQ = make(map[string][]string)
	s.Port = "localhost:8081" // Use different port for testing
	s.Stop = make(chan bool, 1)

	if s.Port != "localhost:8081" {
		t.Errorf("Expected Port to be 'localhost:8081', got %s", s.Port)
	}

	if s.MessageQ == nil {
		t.Errorf("Expected MessageQ to be initialized")
	}

	if s.Stop == nil {
		t.Errorf("Expected Stop channel to be initialized")
	}
}

func TestServerRegistration(t *testing.T) {
	s := new(messaging.ServerService)
	s.MessageQ = make(map[string][]string)
	s.Users = []string{}

	var reply string
	err := s.Register("testuser", &reply)
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}

	if len(s.Users) != 1 || s.Users[0] != "testuser" {
		t.Errorf("Expected Users to contain 'testuser', got %v", s.Users)
	}

	if _, exists := s.MessageQ["testuser"]; !exists {
		t.Errorf("Expected MessageQ to have an entry for 'testuser'")
	}
}

func TestServerGetUsers(t *testing.T) {
	s := new(messaging.ServerService)
	s.Users = []string{"user1", "user2", "user3"}

	var reply []string
	err := s.GetUsers("", &reply)
	if err != nil {
		t.Errorf("GetUsers failed: %v", err)
	}

	expectedLength := 4
	if len(reply) != expectedLength {
		t.Errorf("Expected reply to have %d elements, got %d", expectedLength, len(reply))
	}
}

func TestServerSendMessage(t *testing.T) {
	s := new(messaging.ServerService)
	s.MessageQ = make(map[string][]string)
	s.MessageQ["recipient"] = []string{}

	msg := messaging.Message{
		From:    "sender",
		Target:  "recipient",
		Content: "Hello!",
	}

	var reply messaging.NoReply
	err := s.SendMessage(msg, &reply)
	if err != nil {
		t.Errorf("SendMessage failed: %v", err)
	}

	if len(s.MessageQ["recipient"]) != 1 {
		t.Errorf("Expected recipient's message queue to have 1 message, got %d", len(s.MessageQ["recipient"]))
	}
}

func TestServerCreateRoom(t *testing.T) {
	s := new(messaging.ServerService)
	s.Rooms = []messaging.Room{}

	room := messaging.Room{
		Name:  "testroom",
		Owner: "testowner",
		Users: []string{"testowner"},
	}

	var reply string
	err := s.CreateRoom(room, &reply)
	if err != nil {
		t.Errorf("CreateRoom failed: %v", err)
	}

	if len(s.Rooms) != 1 {
		t.Errorf("Expected Rooms to have 1 element, got %d", len(s.Rooms))
	}

	if s.Rooms[0].Name != "testroom" || s.Rooms[0].Owner != "testowner" {
		t.Errorf("Room not created correctly")
	}
}

func TestServerJoinRoom(t *testing.T) {
	s := new(messaging.ServerService)
	s.Rooms = []messaging.Room{
		{
			Name:  "testroom",
			Owner: "testowner",
			Users: []string{"testowner"},
		},
	}

	room := messaging.Room{
		Name:        "testroom",
		UserJoining: "testuser",
	}

	var reply string
	err := s.JoinRoom(room, &reply)
	if err != nil {
		t.Errorf("JoinRoom failed: %v", err)
	}

	found := false
	for _, user := range s.Rooms[0].Users {
		if user == "testuser" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected user 'testuser' to be added to room")
	}
}

func TestStartServerAndStop(t *testing.T) {
	s := new(messaging.ServerService)
	s.MessageQ = make(map[string][]string)
	s.Port = "localhost:8082"
	s.Stop = make(chan bool, 1)

	messaging.StartServer(s)

	time.Sleep(100 * time.Millisecond)

	var reply messaging.NoReply
	err := s.StopServer("", &reply)
	if err != nil {
		t.Errorf("StopServer failed: %v", err)
	}

	select {
	case <-s.Stop:
	default:
		t.Errorf("Expected stop signal to be sent")
	}
}
