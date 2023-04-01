package cmd

import (
	"log"
	"os"
	"student-app/internal/cmd/migrate"
	"student-app/internal/cmd/serve"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"student-app/internal/config"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg := config.New()
	var logger *zap.Logger
	var err error
	if cfg.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatal(err)
	}
	//nolint: exhaustruct
	root := &cobra.Command{
		Use:   "students",
		Short: "sample application for manage students",
	}

	root.AddCommand(serve.New(logger, cfg))
	root.AddCommand(migrate.New(logger, cfg))

	if err := root.Execute(); err != nil {
		logger.Error("failed to execute root command", zap.Error(err))
		os.Exit(ExitFailure)
	}
}
