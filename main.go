package main

import (
	"github.com/charmbracelet/log"
)

func main() {
    c, err := config()
    if err != nil {
        log.Fatal("Failed to read config", "error", err)
    }

	if c[CONF_RUN_SERVER].(bool) {
		server(&ServerConfig{
			CaCertPath: c[CONF_CA_CERT].(string),
			CertFile:   c[CONF_CERT_FILE].(string),
			KeyFile:    c[CONF_KEY_FILE].(string),
			Port:       uint16(c[CONF_PORT].(int)),
			Server:     c[CONF_LISTEN].(string),
		})
		return
	}

	client(&ClientConfig{
        CaCertPath: c[CONF_CA_CERT].(string),
        CertFile:   c[CONF_CERT_FILE].(string),
        KeyFile:    c[CONF_KEY_FILE].(string),
        Port:       uint16(c[CONF_PORT].(int)),
        Server:     c[CONF_SERVER].(string),
    })
}
