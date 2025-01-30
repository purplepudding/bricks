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

func Test() error {
	mg.Deps()

	return sh.RunV("go", "test", "./...", "-skip", "Integration")
}

func IntegrationTest(svc string) error {
	mg.Deps()

	return sh.RunV("go", "test", "./"+svc+"/...", "-run", "Integration")
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

func New(template string) error {
	mg.Deps() //TODO go install scaffold

	return sh.RunV("scaffold", "new", template)
}
