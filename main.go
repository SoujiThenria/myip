package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

func main() {
    c, err := config()
    if err != nil {
        log.Fatal("Failed to read config", "error", err)
    }

    app := &cli.App {
        Name: "myip",
        Usage: "The client application",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name: "ca",
                Value: c[CONF_CA_CERT].(string),
                Usage: "Path to the root CA file",
            },
            &cli.StringFlag{
                Name: "crt",
                Value: c[CONF_CERT_FILE].(string),
                Usage: "Path to the server certificate",
            },
            &cli.StringFlag{
                Name: "key",
                Value: c[CONF_KEY_FILE].(string),
                Usage: "Path to the server key",
            },
            &cli.IntFlag{
                Name: "port",
                Value: c[CONF_PORT].(int),
                Usage: "The port the server listens on",
            },
            &cli.StringFlag{
                Name: "server",
                Value: c[CONF_SERVER].(string),
                Usage: "The IP or DNS of the server",
            },
        },
        Action: func(cCtx *cli.Context) error {
            client(&ClientConfig{
                CaCertPath: cCtx.String("ca"), 
                CertFile:   cCtx.String("crt"), 
                KeyFile:    cCtx.String("key"), 
                Port:       uint16(cCtx.Int("port")), 
                Server:     cCtx.String("server"), 
            })
            return nil
        },

        Commands: []*cli.Command{
            {
                Name: "server",
                Usage: "Start the server",
                Flags: []cli.Flag{
                    &cli.StringFlag{
                        Name: "ca",
                        Value: c[CONF_CA_CERT].(string),
                        Usage: "Path to the root CA file",
                    },
                    &cli.StringFlag{
                        Name: "crt",
                        Value: c[CONF_CERT_FILE].(string),
                        Usage: "Path to the server certificate",
                    },
                    &cli.StringFlag{
                        Name: "key",
                        Value: c[CONF_KEY_FILE].(string),
                        Usage: "Path to the server key",
                    },
                    &cli.IntFlag{
                        Name: "port",
                        Value: c[CONF_PORT].(int),
                        Usage: "The port the server listens on",
                    },
                    &cli.StringFlag{
                        Name: "listen",
                        Value: c[CONF_LISTEN].(string),
                        Usage: "The IP the server listen on",
                    },
                },
                Action: func(cCtx *cli.Context) error {
                    server(&ServerConfig{
                        CaCertPath: cCtx.String("ca"), 
                        CertFile:   cCtx.String("crt"), 
                        KeyFile:    cCtx.String("key"), 
                        Port:       uint16(cCtx.Int("port")), 
                        Server:     cCtx.String("listen"), 
                    })
                    return nil
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal("Failed to execute", "error", err)
    }
}
