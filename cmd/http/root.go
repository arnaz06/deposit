package main

import (
	"context"
	"log"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/arnaz06/deposit"
	"github.com/arnaz06/deposit/repository"
	"github.com/arnaz06/deposit/service"
	"github.com/lovoo/goka"
	"github.com/spf13/cobra"
)

var (
	depositService service.DepositService
	depositRepo    repository.DepositRepository

	topic goka.Stream = "deposit"
	group goka.Group  = "depisit-group"
)
var rootCmd = &cobra.Command{
	Use:   "deposit",
	Short: "application for managing deposit",
}

func init() {
	cobra.OnInitialize(initApp)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func initApp() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	view, err := goka.NewView([]string{os.Getenv("BROKER_URL")}, goka.GroupTable(group), new(deposit.DepositDecoder), goka.WithViewAutoReconnect())
	if err != nil {
		log.Fatal(err)
	}
	go view.Run(context.Background())

	emitter, err := goka.NewEmitter([]string{os.Getenv("BROKER_URL")}, topic, new(deposit.DepositDecoder))
	if err != nil {
		log.Fatal(err)
	}
	depositRepo = repository.NewDepositRepository(view, emitter)
	depositService = service.NewDepositService(depositRepo)
}
