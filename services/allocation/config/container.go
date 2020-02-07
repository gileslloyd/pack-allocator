//+build wireinject

package config

import (
	"github.com/gileslloyd/gs-allocation-service/internal"
	"github.com/gileslloyd/gs-allocation-service/internal/domain/allocation"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure/controller"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure/microrepo"
	"github.com/google/wire"
)

func CreateApp() internal.App {
	panic(wire.Build(
		GetRoutes,
		infrastructure.NewRouter,
		internal.NewApp,
	))
}

func CreateAllocationController() controller.Allocation {
	panic(wire.Build(
		microrepo.NewMicroPackRepo,
		allocation.NewPackAllocationRule,
		allocation.NewAllocationService,
		controller.NewAllocationController,
	))
}
