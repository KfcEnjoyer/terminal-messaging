package main

import (
	"fmt"
	"terminal-messaging/internal/messaging"
)

func main() {
	user := messaging.UserService{}

	user.Client, _ = user.GetConnection()
	user.GetUsernameFromUser()
	go user.ReadMessages()
	messaging.MainMenu(&user)
	fmt.Println("Stopped connection")

}
