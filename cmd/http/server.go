package main

import (
	handler "github.com/arnaz06/deposit/internal/http"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const port = ":8080"

var serverCmd = &cobra.Command{
	Use:   "http",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		handler.AddDepositHandler(e, depositService)

		log.Info("Starting HTTP server at ", port)
		err := e.Start(port)
		if err != nil {
			log.Fatalf("Failed to start server: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
