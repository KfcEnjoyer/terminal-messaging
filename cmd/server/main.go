package main

import (
	"fmt"
	"terminal-messaging/internal/messaging"
)

func main() {
	s := new(messaging.ServerService)
	s.MessageQ = make(map[string][]string)
	s.Port = "localhost:8080"
	s.Stop = make(chan bool, 1)
	messaging.StartServer(s)
	fmt.Println("Started server on " + s.Port)
	<-s.Stop
}
