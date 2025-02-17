package messaging

type Room struct {
	Name  string
	Owner string
	Users []string
	// Add a field for the user performing the action (if applicable)
	UserJoining string // this field would only be used for join requests
}
