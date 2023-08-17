package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"project/src/pkg/utils"
	cfg "project/src/services/main/configs"
	controller "project/src/services/main/internal/controller/main_service"
	auth "project/src/services/main/internal/handler/auth"
	httpHandler "project/src/services/main/internal/handler/http"
	router "project/src/services/main/internal/handler/http/routes"
	usersRepository "project/src/services/main/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func RunHttpServer(config *cfg.Config) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	u := utils.New()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		utils.FatalResult("Error at set trustedProxies: ", err)
	}

	pasetoAuth, err := auth.NewPasetoFactory(config.API.TokenSecret)
	if err != nil {
		utils.FatalResult("Error at building auth manager: ", err)
	}
	usersRepo, err := usersRepository.NewRepository(config, pasetoAuth, u)
	if err != nil {
		utils.FatalResult("Error while connecting to the database a returning a new repository: ", err)
	}

	ctrl := controller.New(usersRepo)
	h := httpHandler.New(ctrl)
	router.Router(r, h, config, pasetoAuth)

	srv := BuildAndReturnSrv(r, uint32(config.API.Port))

	log.Printf("Http server running at %s:%d", config.API.Domain, config.API.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.FatalResult("Error at set starting server: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	conn := <-ctx.Done()
	log.Println("timeout of 3 seconds. ", conn)
	log.Println("Server exiting")
}

func ShotDownServer(srv *http.Server) error {
	ctx := context.Background()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("Server exiting")
	return nil
}

func BuildAndReturnSrv(r *gin.Engine, port uint32) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	return srv
}

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}
