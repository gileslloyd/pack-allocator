//+build wireinject

package config

import (
	"github.com/gileslloyd/gs-allocation-service/internal"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure"
	"github.com/google/wire"
)

func CreateApp() internal.App {
	panic(wire.Build(
		GetRoutes,
		infrastructure.NewRouter,
		internal.NewApp,
	))
}
