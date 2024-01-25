package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"myip/internal"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type ClientConfig struct {
	RootCA    string `json:"CaCertPath"`
	CLientCrt string `json:"CertFile"`
	ClientKey string `json:"KeyFile"`
	Port      uint16 `json:"Port"`
	Server    string `json:"Server"`
}

func newClientCommand() *cobra.Command {
	confPath := "$HOME/.config/myip/conf.json"
	opts := &ClientConfig{
		RootCA:    "/usr/local/etc/ssl/server.crt",
		CLientCrt: "$HOME/.config/myip/client.crt",
		ClientKey: "$HOME/.config/myip/client.key",
		Port:      443,
		Server:    "localhost",
	}
	clientCmd := &cobra.Command{
		Use:   "myip",
		Short: "The client side of the myip utility",
		Run: func(cmd *cobra.Command, args []string) {
			if err := internal.ReadConfig(confPath, opts); err != nil {
				log.Warn("Failed to read configuration file; use default values", "error", err)
			}
			client(opts)
		},
	}

	clientCmd.Flags().StringVar(&opts.RootCA, "ca", opts.RootCA, "Certificate of the root CA")
	clientCmd.Flags().StringVar(&opts.CLientCrt, "crt", opts.CLientCrt, "Client certificate")
	clientCmd.Flags().StringVar(&opts.ClientKey, "key", opts.ClientKey, "Client key")
	clientCmd.Flags().StringVar(&opts.Server, "listen", opts.Server, "The address the server binds on")
	clientCmd.Flags().Uint16Var(&opts.Port, "port", opts.Port, "The port to use")
	clientCmd.Flags().StringVarP(&confPath, "config", "c", confPath, "The path to the config file")

	return clientCmd
}

func client(c *ClientConfig) {
	caCert, err := os.ReadFile(c.RootCA)
	if err != nil {
		log.Fatal("Failed to read certificate", "error", err)
	}
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Warn("No certificates where parsed")
	}

	cert, err := tls.LoadX509KeyPair(c.CLientCrt, c.ClientKey)
	if err != nil {
		log.Fatal("Failed to load the client certificate and key", "error", err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	resp, err := client.Get("https://" + internal.BuildAddress(c.Server, c.Port))
	if err != nil {
		log.Fatal("Failed to connect to the server", "error", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Warn("Failed to close response body", "error", err)
		}
	}()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body", "error", err)
	}
	fmt.Println(string(respByte))
}
