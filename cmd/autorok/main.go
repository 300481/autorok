package main

import (
	"log"
	"os"

	"github.com/300481/autorok/pkg/cmd/autorok"

	"github.com/urfave/cli"
)

const (
	env_config_url = "AUTOROK_CONFIG_URL"
)

var (
	configUrl string
	app       = cli.NewApp()
)

func init() {
	val, ok := os.LookupEnv(env_config_url)
	if !ok {
		log.Fatalf("Please set environment variable '%s'", env_config_url)
	} else {
		configUrl = val
	}
}

func info() {
	app.Name = "AutoROK"
	app.Usage = "Deployment service for RancherOS based Kubernetes systems"
	app.Author = "Dennis Riemenschneider"
	app.Version = "0.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run in server mode",
			Action: func(c *cli.Context) {
				autorok.NewAutorok(configUrl).Serve()
			},
		},
		{
			Name:    "getrke",
			Aliases: []string{"r"},
			Usage:   "Writes the RKE cluster.yaml to stdout",
			Action: func(c *cli.Context) {
				autorok.NewAutorok(configUrl).GetRKE()
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
