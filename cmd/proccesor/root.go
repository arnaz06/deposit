package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/arnaz06/deposit"
	"github.com/arnaz06/deposit/internal/proccesor"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/spf13/cobra"
)

var (
	topic goka.Stream = "deposit"
	group goka.Group  = "depisit-group"

	tmc *goka.TopicManagerConfig
)

var rootCmd = &cobra.Command{
	Use:   "proccesor",
	Short: "Start proccesor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	tmc = goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1
	cobra.OnInitialize(initApp)
}

func runProcessor(initialized chan struct{}) {
	fmt.Println("processor started!!!!!")

	g := goka.DefineGroup(group,
		goka.Input(topic, new(codec.String), proccesor.Proccess),
		goka.Persist(new(deposit.DepositDecoder)),
	)
	p, err := goka.NewProcessor([]string{os.Getenv("BROKER_URL")},
		g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)
	if err != nil {
		panic(err)
	}

	close(initialized)

	p.Run(context.Background())
}

func initApp() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	tm, err := goka.NewTopicManager([]string{os.Getenv("BROKER_URL")}, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	defer tm.Close()
	err = tm.EnsureStreamExists(string(topic), 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", topic, err)
	}
	initialized := make(chan struct{})

	for {
		runProcessor(initialized)
	}
}
