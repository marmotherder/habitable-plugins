package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func NewPluginObject() interface{} {
	return Exec{}
}

type Exec struct{}

func (c Exec) PluginObject() interface{} {
	return c
}

func (c Exec) Run(directory string, command string, args ...string) (int, string, string) {
	cmd := exec.Command(command, args...)

	cmd.Dir = directory

	stdOut, err := cmd.Output()
	if err != nil {
		exitCode := 0
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			}
		}

		if stdOut != nil {
			return exitCode, "", fmt.Sprintf("%s\n\n%s", err.Error(), string(stdOut))
		}
		return exitCode, "", err.Error()
	}

	return 0, string(stdOut), ""
}
