package server

import (
	"log"
	"os"
	"os/signal"
	"rate/internal/server/handlers"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

type App struct {
	router       *gin.Engine
	rateHandler  handlers.RateHandler
	subscHandler handlers.SubscriptionHandler
}

func NewApp(subscriptionHandler handlers.SubscriptionHandler, rateHandler handlers.RateHandler) *App {

	app := &App{
		router:       gin.Default(),
		subscHandler: subscriptionHandler,
		rateHandler:  rateHandler,
	}
	app.setUpRoutes()
	return app

}

func (app *App) setUpRoutes() {

	app.router.POST("/subscribe", app.subscHandler.Subscribe)
	app.router.GET("/rate", app.rateHandler.Get)

}

func (app *App) Run() {
	port := viper.GetString("SERVER_PORT")
	if err := app.router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server %v", err.Error())
	}
}

func (app *App) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	os.Exit(0)
}
