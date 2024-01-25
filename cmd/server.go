package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"myip/internal"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type ServerOptions struct {
	RootCA    string `json:"RootCA"`
	ServerCrt string `json:"ServerCrt"`
	ServerKey string `json:"ServerKey"`
	Port      uint16 `json:"Port"`
	Listen    string `json:"Listen"`
}

func newServerCommand() *cobra.Command {
	confPath := "/usr/local/etc/myip/conf.json"
	opts := &ServerOptions{
		RootCA:    "/usr/local/etc/ssl/server.crt",
		ServerCrt: "/usr/local/etc/myip/server.crt",
		ServerKey: "/usr/local/etc/myip/server.key",
		Port:      443,
		Listen:    "0.0.0.0",
	}
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "The server side of the myip utility",
		Run: func(cmd *cobra.Command, args []string) {
			if err := internal.ReadConfig(confPath, opts); err != nil {
				log.Debug("No config file found", "error", err)
			}
			server(opts)
		},
	}

	serverCmd.Flags().StringVar(&opts.RootCA, "ca", opts.RootCA, "Certificate of the root CA")
	serverCmd.Flags().StringVar(&opts.ServerCrt, "crt", opts.ServerCrt, "Client certificate")
	serverCmd.Flags().StringVar(&opts.ServerKey, "key", opts.ServerKey, "Client key")
	serverCmd.Flags().StringVar(&opts.Listen, "listen", opts.Listen, "The address the server binds on")
	serverCmd.Flags().Uint16Var(&opts.Port, "port", opts.Port, "The port to use")
	serverCmd.Flags().StringVarP(&confPath, "config", "c", confPath, "The path to the config file")

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
		Addr:      internal.BuildAddress(c.Listen, c.Port),
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		clientAddress := r.RemoteAddr
		log.Info("New connection", "address", clientAddress, "CN", r.TLS.PeerCertificates[0].Subject)
		w.Write([]byte(clientAddress[:strings.LastIndex(clientAddress, ":")]))
	})

	log.Info("Start server", "server", c.Listen, "port", c.Port)
	if err := server.ListenAndServeTLS(c.ServerCrt, c.ServerKey); err != nil {
		log.Fatal("Failed to start the server", "error", err)
	}
}
