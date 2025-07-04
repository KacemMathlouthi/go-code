package tools

import "os/exec"

func Shell(command string) (string, error) {
	out, err := exec.Command("sh", "-c", command).CombinedOutput()
	return string(out), err
}
