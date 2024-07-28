package main

import (
	"fmt"

	"github.com/rayelissonl/goapreapi/config"
	"github.com/rayelissonl/goapreapi/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		fmt.Println("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
