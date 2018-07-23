// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Build() error {
	fmt.Println("Building...")
	env := map[string]string{
		"GOARCH":"amd64",
		"GOOS":"linux",
	}
	hash, err := sh.OutputWith(env,"go","build","withings.go")
	fmt.Println(hash)
	return err
}

func LocalRun() error{
	//sam local invoke -e event.json -t template.yaml --env-vars sam_env.json
	fmt.Println("Start...")
	cmd :=[]string{"local","invoke",
		"-e","event.json",
		"-t","template.yaml",
		"--env-vars","sam_env.json"}

	hash,err :=sh.Output("sam",cmd...)
	fmt.Println(hash)
	return err
}