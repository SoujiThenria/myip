package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type ServerOptions struct {
    RootCA string
    ClientCrt string
    ClientKey string
    Port uint16
    Listen string
}

func newServerCommand() *cobra.Command {
    opts := new(ServerOptions)
    serverCmd := &cobra.Command{
        Use:   "server",
        Short: "The server side of the myip utility",
        Run: func(cmd *cobra.Command, args []string) {
            log.Info("Start the server")
            server(opts)
        },
    }

    serverCmd.Flags().StringVar(&opts.RootCA, "ca", "/etc/ssl/server.crt", "Certificate of the root CA")
    serverCmd.Flags().StringVar(&opts.ClientCrt, "crt", "/usr/local/etc/myip/server.crt", "Client certificate")
    serverCmd.Flags().StringVar(&opts.ClientKey, "key", "/usr/local/etc/myip/server.key", "Client key")
    serverCmd.Flags().StringVar(&opts.Listen, "listen", "0.0.0.0", "The address the server binds on")
    serverCmd.Flags().Uint16Var(&opts.Port, "port", 443, "The port to use")

    return serverCmd
}

func server(c *ServerOptions) {
    log.Info(c)
}
