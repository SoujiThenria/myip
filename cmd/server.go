package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type ServerOptions struct {
    RootCA string
    ServerCrt string
    ServerKey string
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
    serverCmd.Flags().StringVar(&opts.ServerCrt, "crt", "/usr/local/etc/myip/server.crt", "Client certificate")
    serverCmd.Flags().StringVar(&opts.ServerKey, "key", "/usr/local/etc/myip/server.key", "Client key")
    serverCmd.Flags().StringVar(&opts.Listen, "listen", "0.0.0.0", "The address the server binds on")
    serverCmd.Flags().Uint16Var(&opts.Port, "port", 443, "The port to use")

    return serverCmd
}

func server(c *ServerOptions) {
	caCert, err := os.ReadFile(c.RootCA)
	if err != nil {
		log.Fatal("Failed to read certificate", "error", err)
	}
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Warn("No certificates where parsed")
	}

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	server := &http.Server{
		Addr:      buildAddress(c.Listen, c.Port),
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		clientAddress := r.RemoteAddr
		log.Info("New connection", "client", clientAddress)
		w.Write([]byte(clientAddress[:strings.LastIndex(clientAddress, ":")]))
	})

	log.Info("Start server", "server", c.Listen, "port", c.Port)
	if err := server.ListenAndServeTLS(c.ServerCrt, c.ServerKey); err != nil {
		log.Fatal("Failed to start the server", "error", err)
	}
}
