package main

import (
	"fmt"
	"time"

	"project/src/pkg/utils"
	controller "project/src/services/main/internal/controller/main_service"
	httpHandler "project/src/services/main/internal/handler/http"
	router "project/src/services/main/internal/handler/http/routes"
	memory "project/src/services/main/internal/repository/memory"

	"github.com/gin-gonic/gin"
)

func RunHttpServer() {
	fmt.Println("Running http server")
	r := gin.Default()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		utils.FatalResult("Error at set trustedProxies: ", err)
	}
	memory := memory.New()
	ctrl := controller.New(memory)
	h := httpHandler.New(ctrl)
	router.Router(r, h)
	err = r.Run(":4000")
	if err != nil {
		utils.FatalResult("Error at set starting server: ", err)
	}
}

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}
