package messaging

type Room struct {
	Name        string
	Owner       string
	Users       []string
	UserJoining string
}
