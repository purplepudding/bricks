//go:build mage

package main

import (
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Gen() error {
	mg.Deps()

	//TODO make this lazy and only generate the files that need changing
	return sh.RunV("go", "generate", "-v", "-x", "./...")
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

func Deploy(env string) error {
	mg.Deps()

	c := "kustomize build deploy/k8s/" + env + "| kubectl apply -f -"
	cmd := exec.Command("bash", "-c", c)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
