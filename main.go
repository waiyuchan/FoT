package main

import (
	"FoT/configs/init"
	"FoT/routers/router"
)

func main() {
	config := init.initConfig()
	router := router.registerRouter()
	router.Run(config.Get("settings.application.port"))
}

