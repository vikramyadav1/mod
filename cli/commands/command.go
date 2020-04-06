package commands

// Command interface to execute commands
type Command interface {
	Execute() error
}
