// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package config

import (
	"github.com/gileslloyd/gs-allocation-service/internal"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure"
)

// Injectors from container.go:

func CreateApp() internal.App {
	v := GetRoutes()
	router := infrastructure.NewRouter(v)
	app := internal.NewApp(router)
	return app
}
