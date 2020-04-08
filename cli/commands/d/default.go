package d

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vikramyadav1/mod/cli/config"
)

type def struct {
	Args []string
}

// Execute : Implements command interface
func (df def) Run() error {
	imageName := df.Args[0]
	c := config.NewModConfig().D

	volumeOptions := strings.Join(c.Defaults.Volumes, "-v ")
	portOptions := strings.Join(c.Defaults.Ports, "-p")
	cmdargs := []string{"run", "-it", "-d", "-v", volumeOptions, "-p", portOptions, imageName, c.Defaults.Command}

	cmd := exec.Command("docker", cmdargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	return nil
}
