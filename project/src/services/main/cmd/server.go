package main

import (
	"fmt"
	"time"

	"project/src/pkg/utils"
	cfg "project/src/services/main/configs"
	controller "project/src/services/main/internal/controller/main_service"
	httpHandler "project/src/services/main/internal/handler/http"
	router "project/src/services/main/internal/handler/http/routes"
	memory "project/src/services/main/internal/repository/memory"

	"github.com/gin-gonic/gin"
)

func RunHttpServer() {
	config := GetConfig()
	r := gin.Default()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{config.API.Domain})
	if err != nil {
		utils.FatalResult("Error at set trustedProxies: ", err)
	}
	memory := memory.New()
	ctrl := controller.New(memory)
	h := httpHandler.New(ctrl)
	router.Router(r, h)
	err = r.Run(fmt.Sprintf(":%d", config.API.Port))
	if err != nil {
		utils.FatalResult("Error at set starting server: ", err)
	}
	fmt.Println("Running http server")
}

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}

// * Get a pointer to config object with all
// * needed variables from yaml file
func GetConfig() *cfg.Config {
	cfg, err := cfg.LoadConfig()
	if err != nil {
		utils.FatalResult("Could not load config: ", err)
	}
	return cfg
}
