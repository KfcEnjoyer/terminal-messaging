package messaging_test

import (
	"terminal-messaging/internal/messaging"
	"testing"
)

func TestMessageStructure(t *testing.T) {
	msg := messaging.Message{
		From:    "sender",
		Target:  "recipient",
		Content: "Hello, world!",
	}

	if msg.From != "sender" {
		t.Errorf("Expected From to be 'sender', got %s", msg.From)
	}

	if msg.Target != "recipient" {
		t.Errorf("Expected Target to be 'recipient', got %s", msg.Target)
	}

	if msg.Content != "Hello, world!" {
		t.Errorf("Expected Content to be 'Hello, world!', got %s", msg.Content)
	}
}

func TestRoomStructure(t *testing.T) {
	room := messaging.Room{
		Name:        "testroom",
		Owner:       "testowner",
		Users:       []string{"user1", "user2"},
		UserJoining: "newuser",
	}

	if room.Name != "testroom" {
		t.Errorf("Expected Name to be 'testroom', got %s", room.Name)
	}

	if room.Owner != "testowner" {
		t.Errorf("Expected Owner to be 'testowner', got %s", room.Owner)
	}

	if len(room.Users) != 2 {
		t.Errorf("Expected Users to have 2 elements, got %d", len(room.Users))
	}

	if room.UserJoining != "newuser" {
		t.Errorf("Expected UserJoining to be 'newuser', got %s", room.UserJoining)
	}
}
