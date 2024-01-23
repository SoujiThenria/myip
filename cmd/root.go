package cmd

import (
    "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type ClientOptions struct {
}

var rootCmd = &cobra.Command{
  Use:   "myip",
  Short: "The client side of the myip utility",
  Version: "0.1-beta",
  Run: func(cmd *cobra.Command, args []string) {
      log.Info("Start client")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
      log.Fatal("Failed to execute", "error", err)
  }
}
