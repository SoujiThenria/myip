package main

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	CONF_CA_CERT    = "cacert"
	CONF_CERT_FILE  = "certfile"
	CONF_KEY_FILE   = "keyfile"
	CONF_LISTEN     = "listen"
	CONF_PORT       = "port"
	CONF_SERVER     = "server"
	CONF_RUN_SERVER = "runserver"
	CONF_RUN_CLIENT = "runclient"
)

func config() (config map[string]any, err error) {
	viper.SetDefault(CONF_CA_CERT, "/etc/ssl/rootCA.crt")
	viper.SetDefault(CONF_CERT_FILE, "/etc/ssl/myip.crt")
	viper.SetDefault(CONF_KEY_FILE, "/etc/ssl/myip.key")
	viper.SetDefault(CONF_LISTEN, "0.0.0.0")
	viper.SetDefault(CONF_PORT, 443)
	viper.SetDefault(CONF_SERVER, "localhost")
	viper.SetDefault(CONF_RUN_SERVER, false)
	viper.SetDefault(CONF_RUN_CLIENT, true)

	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/myip")
	viper.AddConfigPath("/usr/local/etc/myip/")
	viper.AddConfigPath("/etc/myip/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("No configuration found, using defaults")
			// log.Info("Write configuration")
			// // TODO: does not create directory
			// // TODO: does not resolve $HOME
			// err = viper.SafeWriteConfigAs("/home/souji/.config/myip/conf.toml")
			// if err != nil {
			//     log.Fatal("Failed to write config", "error", err)
			// }
		} else {
            return nil, err
		}
	}

    pflag.String(CONF_CA_CERT, "", "Path to the root CA certificate")
    pflag.String(CONF_CERT_FILE, "", "Path to the certificate")
    pflag.String(CONF_KEY_FILE, "", "Path to the key file")
    pflag.String(CONF_LISTEN, "", "Address the server listens on")
    pflag.Int16P(CONF_PORT, "p", 9001, "The port to use")
    pflag.String(CONF_SERVER, "", "The server the client connects to")
    pflag.BoolP(CONF_RUN_SERVER, "s", false, "Run the server")
    pflag.BoolP(CONF_RUN_CLIENT, "c", false, "Run the client")
    pflag.Parse()
    viper.BindPFlags(pflag.CommandLine)

    // Yeah, thats an ugly hack
    viper.Set(CONF_PORT, viper.GetInt(CONF_PORT))

	return viper.AllSettings(), nil
}
