package common

import (
	"os"
)

type Config struct {
	Environment  string
	Debug        bool
	DbHost       string
	DbUser       string
	DbPassword   string
	DbName       string
	DbPort       string
	FlickrApiKey string
}

var AppConfig Config

func init() {
	AppConfig.Environment = os.Getenv("ENVIRONMENT")
	AppConfig.Debug = "true" == os.Getenv("DEBUG")
	AppConfig.DbHost = os.Getenv("MYSQL_HOST")
	AppConfig.DbUser = os.Getenv("MYSQL_USER")
	AppConfig.DbPassword = os.Getenv("MYSQL_PASSWORD")
	AppConfig.DbName = os.Getenv("MYSQL_DATABASE")
	AppConfig.DbPort = os.Getenv("MYSQL_PORT")

	AppConfig.FlickrApiKey = os.Getenv("FLICKR_API_KEY")
}
