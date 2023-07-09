package main

import (
	"fmt"
	"time"
)

func RunHttpServer() {
	fmt.Println("Running http server")
}

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}
