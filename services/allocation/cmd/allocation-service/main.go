package main

import (
	"github.com/gileslloyd/gs-allocation-service/config"
)

func main() {
	app := config.CreateApp()

	app.Start()
}
