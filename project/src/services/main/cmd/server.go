package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RunHttpServer() {
	fmt.Println("Running http server")
	r := gin.Default()

	r.GET("/ping", ping)
	r.Run(":4000")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}
