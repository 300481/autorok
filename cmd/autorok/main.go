package main

import (
	"log"
	"os"

	"github.com/300481/autorok/pkg/cmd/autorok"
)

const (
	env_config_url = "AUTOROK_CONFIG_URL"
)

var (
	configUrl string
)

func init() {
	val, ok := os.LookupEnv(env_config_url)
	if !ok {
		log.Fatalf("Please set environment variable '%s'", env_config_url)
	} else {
		configUrl = val
	}
}

func main() {
	autorok.NewAutorok(configUrl).Execute()
}
