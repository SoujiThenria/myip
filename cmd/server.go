package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "The server side of the myip utility",
  Run: func(cmd *cobra.Command, args []string) {
      log.Info("Start the server")
  },
}
