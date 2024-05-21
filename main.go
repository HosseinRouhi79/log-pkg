package main

import (
	"github.com/HosseinRouhi79/log-pkg/config"
	"github.com/HosseinRouhi79/log-pkg/logging"
)


func main(){
	cfg := config.LogConfig()
	cfg.Logger = "zerolog"
	logger := logging.NewLogger(&cfg)

	logger.Infof("%s", "logger")
}		