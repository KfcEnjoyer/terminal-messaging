package main

import (
	"flag"
	"fmt"
	"terminal-messaging/internal/messaging"
)

func main() {
	portFlag := flag.String("port", "8080", "Port number for the server to listen on")
	hostFlag := flag.String("host", "localhost", "Host address for the server to bind to")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *hostFlag, *portFlag)

	s := new(messaging.ServerService)
	s.MessageQ = make(map[string][]string)
	s.Port = address
	s.Stop = make(chan bool, 1)
	messaging.StartServer(s)
	fmt.Println("Started server on " + s.Port)
	<-s.Stop
}
