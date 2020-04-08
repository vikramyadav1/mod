package d

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/vikramyadav1/essentials/slice"
)

type rmi struct {
	Args []string
}

func (r rmi) Run() error {
	cmd := exec.Command("bash", "-c", "docker images | tr -s ' ' | cut -d ' ' -f 3")
	force := Contains(r.Args, "-f")

	if bs, err := cmd.Output(); err != nil {
		fmt.Println(err)
	} else {
		output := string(bs)
		containers := strings.Split(output, "\n")

		for _, imageID := range containers[1 : len(containers)-1] {
			removeImage(imageID, force)
		}
	}

	return nil
}

func removeImage(imageID string, force bool) {
	var rmcommand *exec.Cmd

	if force == true {
		rmcommand = exec.Command("docker", "rmi", "-f", imageID)
	} else {
		rmcommand = exec.Command("docker", "rmi", imageID)
	}

	rmcommand.Stdout = os.Stdout
	rmcommand.Stderr = os.Stderr

	if err := rmcommand.Run(); err != nil {
		fmt.Println(err)
	}
}
