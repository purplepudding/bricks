//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Gen() error {
	mg.Deps() //TODO add go install for all tools

	//TODO make this lazy and only generate the files that need changing
	return sh.RunV("go", "generate", "./...")
}

func Dev(svc string) error {
	//TODO validate path?
	if err := os.Chdir(svc); err != nil {
		return err
	}
	defer func() {
		_ = os.Chdir("..")
	}()

	return sh.RunV("skaffold", "dev", "--port-forward")
}
