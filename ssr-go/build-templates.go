package main

import (
	"os"
	"os/exec"
)

func main() {
	// Run the templ build command
	execPath, err := exec.LookPath("templ")

	if err != nil {
		panic(err)
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Run the command
	cmd := exec.Command(execPath, "generate")
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
