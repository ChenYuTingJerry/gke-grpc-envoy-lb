package main

import (
	"fmt"

	app "http-service/internal/injector"
)

func main() {
	fmt.Println("main")
	app.Run()
}
