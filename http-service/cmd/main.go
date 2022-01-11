package main

import (
	"fmt"

	app "github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/internal/injector"
)

func main() {
	fmt.Println("main")
	app.Run()
}
