package main

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type ServerConfig struct {
	CaCertPath string
	CertFile   string
	KeyFile    string
	Port       uint16
	Server     string
}

func server(c *ServerConfig) {
	caCert, err := os.ReadFile(c.CaCertPath)
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
		Addr:      buildAddress(c.Server, c.Port),
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		clientAddress := r.RemoteAddr
		log.Info("New connection", "client", clientAddress)
		w.Write([]byte(clientAddress[:strings.LastIndex(clientAddress, ":")]))
	})

	log.Info("Start server", "server", c.Server, "port", c.Port)
	if err := server.ListenAndServeTLS(c.CertFile, c.KeyFile); err != nil {
		log.Fatal("Failed to start the server", "error", err)
	}
}

func buildAddress(server string, port uint16) string {
	return server + ":" + strconv.Itoa(int(port))
}
