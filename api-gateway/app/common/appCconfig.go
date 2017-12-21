package common

import (
	"os"
)

type Config struct {
	Environment string
	Debug       bool
	BackendHost string
	BackendPort string
}

var AppConfig Config

func init() {
	AppConfig.Environment = os.Getenv("ENVIRONMENT")
	AppConfig.Debug = "true" == os.Getenv("DEBUG")
	AppConfig.BackendHost = os.Getenv("BACKEND_HOST")
	AppConfig.BackendPort = os.Getenv("BACKEND_PORT")
}
