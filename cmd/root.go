package cmd

import "github.com/charmbracelet/log"

func Execute() {
	rootCmd := newClientCommand()
	rootCmd.Version = "1.0.0"

	rootCmd.AddCommand(newServerCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Failed to execute", "error", err)
	}
}
