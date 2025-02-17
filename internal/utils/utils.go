package utils

import (
	"errors"
	"strings"
)

func ValidateUsername(username string, users []string, length int) error {
	if strings.TrimSpace(username) == "" {
		return errors.New("username cannot be blank")
	}
	if len(strings.TrimSpace(username)) < length {
		return errors.New("username cannot be less than " + string(length) + " characters")
	}

	for _, user := range users {
		if username == user {
			return errors.New("this user already exists")
		}
	}

	return nil
}

func ValidateParams(params []string) error {
	if params[0] == "send" && params[1] == strings.TrimSpace("") {
		return errors.New("target username cannot be blank")
	} else if params[0] == "global" && params[1] == strings.TrimSpace("") {
		return errors.New("you cannot send a blank message")
	} else if params[0] == "create" && params[1] == strings.TrimSpace("") {
		return errors.New("room name cannot be blank")
	} else if params[0] == "join" && params[1] == strings.TrimSpace("") {
		return errors.New("room name cannot be blank")
	} else if params[0] == "leave" && params[1] == strings.TrimSpace("") {
		return errors.New("room name cannot be blank")
	} else if params[0] == "sendroom" && params[1] == strings.TrimSpace("") {
		return errors.New("room name cannot be blank")
	} else if params[0] == "showroomus" && params[1] == strings.TrimSpace("") {
		return errors.New("room name cannot be blank")
	}
	return nil
}
