package msg

// ErrorMsg allows passing errors through the Bubble Tea program
type ErrorMsg error

// QuitMsg is sent when the user wants to quit the application
type QuitMsg struct{}

// LoginSuccessMsg is sent when login is successful
type LoginSuccessMsg struct {
	Token    string
	UserName string
}

// LoginFailMsg is sent when login fails
type LoginFailMsg error

// SwitchToRegisterMsg switches view to register
type SwitchToRegisterMsg struct{}

// SwitchToLoginMsg switches view to login
type SwitchToLoginMsg struct{}
