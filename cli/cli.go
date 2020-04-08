package cli

import (
	"fmt"
	"os"

	"github.com/vikramyadav1/mod/cli/commands"
	"github.com/vikramyadav1/mod/cli/commands/d"
)

// Start : Start the cli
func Start() {
	f := GFlags{}.Parse()

	if f.Version == true {
		fmt.Println(Version)
	} else {
		runSubcommand()
	}
}

func runSubcommand() error {
	if len(os.Args) < 2 {
		fmt.Println("Missing Command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "d":
		return executeCommand(d.D{})
	default:
		return fmt.Errorf("Error: Unrecognized command. Run with -h for usage")
	}
}

func executeCommand(c commands.Command) error {
	return c.Execute()
}
