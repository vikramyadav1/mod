package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vikramyadav1/mod/cli/config"
)

//D : Spin up docker containers
type D struct {
	ImageName string
}

// Execute : Implements command interface
func (d D) Execute() error {
	c := config.NewModConfig().D

	dockerExecutable, _ := exec.LookPath("docker")

	volumeOptions := strings.Join(c.Defaults.Volumes, "-v ")
	portOptions := strings.Join(c.Defaults.Ports, "-p")
	args := []string{dockerExecutable, "run", "-it", "-d", "-v", volumeOptions, "-p", portOptions, d.ImageName, c.Defaults.Command}

	cmd := &exec.Cmd{
		Path:   dockerExecutable,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	return nil
}
