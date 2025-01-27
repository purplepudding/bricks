package main

import (
	"github.com/purplepudding/foundation/internal/service"
)

func main() {
	s := service.NewService()
	s.Run()
}
