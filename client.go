package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

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
