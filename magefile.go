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
