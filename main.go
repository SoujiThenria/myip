package main

import (
	"flag"

	"github.com/charmbracelet/log"
)

type config struct {
	run_server   bool
	port         int
	CACertPath   string
	CertFilePath string
	KeyFilePath  string
	ListenOn     string
	server       string
}

func main() {
	var c config
	flag.BoolVar(&c.run_server, "s", false, "Run the server")
	flag.IntVar(&c.port, "p", 443, "The port where the comunication between server and client takes place")
	flag.StringVar(&c.CACertPath, "ca", "./cert/root_ca.crt", "The path to the root CA certificate")
	flag.StringVar(&c.CertFilePath, "cert", "./cert/server.crt", "The path to the certificate file")
	flag.StringVar(&c.KeyFilePath, "key", "./cert/server.key", "The path to the key file")
	flag.StringVar(&c.ListenOn, "l", "0.0.0.0", "Where the server listens on")
	flag.StringVar(&c.server, "server", "localhost", "The IP or host name where the client is supposed to connect to.")
	flag.Parse()

	if c.port < 0 || c.port > 0xFFFF {
		log.Fatal("Invalid port, dont be a dum dum...")
	}

	if c.run_server {
		server(&ServerConfig{
			CaCertPath: c.CACertPath,
			CertFile:   c.CertFilePath,
			KeyFile:    c.KeyFilePath,
			Port:       uint16(c.port),
			Server:     c.ListenOn,
		})
		return
	}

	client(&ClientConfig{
		CaCertPath: c.CACertPath,
		CertFile:   c.CertFilePath,
		KeyFile:    c.KeyFilePath,
		Port:       uint16(c.port),
		Server:     c.server,
	})
}
