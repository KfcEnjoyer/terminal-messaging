package main

import (
	"fmt"
	"strings"
	"terminal-messaging/internal/messaging"
)

func main() {
	user := messaging.UserService{}

	user.Client, _ = user.GetConnection()
	fmt.Println("Enter your username:")
	var username string
	fmt.Scanln(&username)
	user.Username = strings.TrimSpace(username)
	user.Register(user.Username)
	go user.ReadMessages()
	messaging.MainMenu(&user)
	fmt.Println("Stopped connection")

}
