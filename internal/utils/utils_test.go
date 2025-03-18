package utils_test

import (
	"terminal-messaging/internal/utils"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		users    []string
		length   int
		wantErr  bool
	}{
		{
			name:     "Valid username",
			username: "testuser",
			users:    []string{"user1", "user2"},
			length:   3,
			wantErr:  false,
		},
		{
			name:     "Empty username",
			username: "",
			users:    []string{"user1", "user2"},
			length:   3,
			wantErr:  true,
		},
		{
			name:     "Short username",
			username: "ab",
			users:    []string{"user1", "user2"},
			length:   3,
			wantErr:  true,
		},
		{
			name:     "Username already exists",
			username: "user1",
			users:    []string{"user1", "user2"},
			length:   3,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateUsername(tt.username, tt.users, tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateParams(t *testing.T) {
	tests := []struct {
		name    string
		params  []string
		wantErr bool
	}{
		{
			name:    "Valid send params",
			params:  []string{"send", "user1", "Hello!"},
			wantErr: false,
		},
		{
			name:    "Invalid send params - empty target",
			params:  []string{"send", ""},
			wantErr: true,
		},
		{
			name:    "Valid global params",
			params:  []string{"global", "Hello everyone!"},
			wantErr: false,
		},
		{
			name:    "Invalid global params - empty message",
			params:  []string{"global", ""},
			wantErr: true,
		},
		{
			name:    "Valid create room params",
			params:  []string{"create", "room1"},
			wantErr: false,
		},
		{
			name:    "Invalid create room params - empty name",
			params:  []string{"create", ""},
			wantErr: true,
		},
		{
			name:    "Valid join room params",
			params:  []string{"join", "room1"},
			wantErr: false,
		},
		{
			name:    "Invalid join room params - empty name",
			params:  []string{"join", ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateParams(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
