package main

import (
	"context"
	"time"

	"github.com/danargh/go-clean-arc/config"
	"github.com/danargh/go-clean-arc/internal/app"
	"github.com/labstack/gommon/log"
)

func main() {
	// load config file
	cfg, err := config.LoadConfig("config")
	if err != nil {
		panic(err)
	}

	// set timezhone
	_, err = time.LoadLocation(cfg.Server.TimeZone)
	if err != nil {
		log.Error("Error settning timezone : ", err)
		return
	}

	// build new app server (call NewApp then call Run
	if err := app.NewApp(context.Background(), cfg).Run(); err != nil {
		log.Print("Unable to start app")
		return
	}

}
