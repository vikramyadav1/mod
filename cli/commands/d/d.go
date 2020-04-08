package d

import (
	"flag"
	"fmt"
	"os"
)

//D : Spin up docker container
type D struct{}

// Execute : Execute under d subcommand
func (d D) Execute() error {
	cmd := flag.NewFlagSet("d", flag.ExitOnError)
	cmd.Parse(os.Args[2:])

	args := cmd.Args()

	switch args[0] {
	case "rmall":
		fmt.Println("removing all docker containers")
		return rmall{Args: args}.Run()
	case "rmi":
		fmt.Println("removing all docker images")
		return rmi{Args: args}.Run()
	default:
		return def{Args: args}.Run()
	}
}
