package main

import (
	"project/src/pkg/utils"
	cfg "project/src/services/main/configs"

	"github.com/gin-gonic/gin"
)

// @Title Http-template-api
// @version 0.1.0
// @description A Http simple go api
// @host localhost:4000
// @BasePath /api/v1
func main() {
	config := GetConfig()
	if config.API.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	RunHttpServer(config)
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
