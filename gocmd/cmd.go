package gocmd

import (
	"os/exec"
	"strings"
)

func ExecCmd(cmdStr string) (res string, err error) {
	args := strings.Split(cmdStr, " ")
	resb,err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", err
	}

	return string(resb), nil
}
