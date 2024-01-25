package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
    c *ClientConfig
    rootCmd = &cobra.Command{
        Use:   "myip",
        Short: "The client side of the myip utility",
        Version: "0.1-beta",
        Run: func(cmd *cobra.Command, args []string) {
            client(c)
        },
    }
)

func Execute() {
    c = new(ClientConfig)

    rootCmd.Flags().StringVar(&c.CaCertPath, "ca", "/etc/ssl/server.crt", "Certificate of the root CA")
    rootCmd.Flags().StringVar(&c.CertFile, "crt", "$HOME/.config/myip/client.crt", "Client certificate")
    rootCmd.Flags().StringVar(&c.KeyFile, "key", "$HOME/.config/myip/client.key", "Client key")
    rootCmd.Flags().StringVar(&c.Server, "server", "localhost", "IP or DNS name of the server")
    rootCmd.Flags().Uint16Var(&c.Port, "port", 443, "The port to use")

    rootCmd.AddCommand(newServerCommand())

  if err := rootCmd.Execute(); err != nil {
      log.Fatal("Failed to execute", "error", err)
  }
}

type ClientConfig struct {
	CaCertPath string
	CertFile   string
	KeyFile    string
	Port       uint16
	Server     string
}

func client(c *ClientConfig) {
	caCert, err := os.ReadFile(c.CaCertPath)
	if err != nil {
		log.Fatal("Failed to read certificate", "error", err)
	}
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		log.Warn("No certificates where parsed")
	}

	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
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

	resp, err := client.Get("https://" + buildAddress(c.Server, c.Port))
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

func buildAddress(server string, port uint16) string {
	return server + ":" + strconv.Itoa(int(port))
}

func getFlag(flag string, err error) string {
    if err != nil {
        log.Fatal(err)
    }
    return flag
}
