package main

import (
	"fmt"

	"http-service/configs"
	app "http-service/internal/injector"
)

func main() {
	fmt.Println("main")
	configs.InitConfigs()
	app.Run()
}
