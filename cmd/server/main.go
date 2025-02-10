package main

import (
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/lib"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/logger"
	"os"
)

const configPath = "config/config.yaml"

func main() {
	conf := config.Load(configPath)
	log := logger.SetUpLogger(conf.ENV)

	storageType, err := lib.ParseStorageType(os.Args)
	if err != nil {
		panic(err)
	}

	log.Debug("set up", "conf", conf, "storageType", storageType)
}
