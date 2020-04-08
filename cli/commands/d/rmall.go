package d

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/vikramyadav1/essentials/slice"
)

type rmall struct {
	Args []string
}

func (r rmall) Run() error {
	cmd := exec.Command("bash", "-c", "docker ps -a | cut -d ' ' -f1")
	force := Contains(r.Args, "-f")

	if bs, err := cmd.Output(); err != nil {
		fmt.Println(err)
	} else {
		output := string(bs)
		containers := strings.Split(output, "\n")

		for _, containerID := range containers[1 : len(containers)-1] {
			removeContainer(containerID, force)
		}
	}

	return nil
}

func removeContainer(containerID string, force bool) {
	var rmcommand *exec.Cmd

	if force == true {
		rmcommand = exec.Command("docker", "rm", "-f", containerID)
	} else {
		rmcommand = exec.Command("docker", "rm", containerID)
	}

	rmcommand.Stdout = os.Stdout
	rmcommand.Stderr = os.Stderr

	if err := rmcommand.Run(); err != nil {
		fmt.Println(err)
	}
}
