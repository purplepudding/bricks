package main

import (
	"github.com/purplepudding/foundation/auth/internal/service"
)

func main() {
	s := service.NewService()
	s.Run()
}
