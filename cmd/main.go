package main

import (
	"log"
	"rate/internal/database"
	"rate/internal/pkg/mailer"
	"rate/internal/pkg/provider"
	"rate/internal/pkg/provider/chain"
	"rate/internal/server"
	"rate/internal/server/handlers"
	"rate/internal/services"
	"sync"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/robfig/cron"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}

	db := database.InitDB()

	database.RunDBMigration()

	subscRepo := database.NewSubscriptionRepo(db)
	subscService := services.NewSubscriptionService(subscRepo)
	subscHandler := handlers.NewSubscriptionHandler(subscService)

	baseRateProviderNode := chain.NewProviderNode(provider.NewRateProvider())
	secRateProviderNode := chain.NewProviderNode(provider.NewSecRateProvider())
	baseRateProviderNode.SetNext(secRateProviderNode)

	rateService := services.NewRateProviderService(baseRateProviderNode)
	rateHandler := handlers.NewRateHandler(rateService)

	smtpServer := mailer.NewSMTPServer(viper.GetString("MAILER_HOST"), viper.GetString("MAILER_PORT"), viper.GetString("MAILER_USERNAME"), viper.GetString("MAILER_PASSWORD"), viper.GetString("MAILER_FROM"))

	mailWG := &sync.WaitGroup{}
	emailSevice := services.NewRateMailerService(smtpServer, subscRepo, baseRateProviderNode, mailWG)

	cronOperator := cron.New()
	err = cronOperator.AddFunc("00 29 12 * * *", emailSevice.SendEmails)
	if err != nil {
		log.Fatal("Failed to configure cron operation:", err)
	}

	cronOperator.Start()

	app := server.NewApp(*subscHandler, *rateHandler)
	log.Println("App struct was created")
	app.Run()
	go app.ListenForShutdown()
}
