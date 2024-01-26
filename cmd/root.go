package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	logLevel = "INFO"
    Version = "development"
)

func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd := newClientCommand()
	rootCmd.Version = Version

	rootCmd.AddCommand(newServerCommand())

	// Global flags
	rootCmd.PersistentFlags().StringVarP(&logLevel, "logLevel", "l", logLevel, "log level")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute", "error", err)
	}
}

func initConfig() {
	l, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("Failed to parse log level", "error", err)
	}
	log.SetLevel(l)
}
