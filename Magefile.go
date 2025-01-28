//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Gen() error {
	mg.Deps()

	//TODO make this lazy and only generate the files that need changing
	return sh.RunV("go", "generate", "./...")
}
