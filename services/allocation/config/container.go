//+build wireinject

package config

import (
	controller2 "github.com/gileslloyd/gs-allocation-service/internal/app/controller"
	"github.com/gileslloyd/gs-allocation-service/internal/domain/allocation"
	"github.com/gileslloyd/gs-allocation-service/pkg/infrastructure/delivery/rpc"
	"github.com/gileslloyd/gs-allocation-service/pkg/infrastructure/microrepo"
	"github.com/google/wire"
)

func CreateAllocationController() controller2.Allocation {
	panic(wire.Build(
		microrepo.NewMicroPackRepo,
		allocation.NewPackAllocationRule,
		allocation.NewAllocationService,
		controller2.NewAllocationController,
	))
}

func CreateMessageListener() rpc.MessageListener {
	panic(wire.Build(
		GetRoutes,
		rpc.NewHandler,
		rpc.NewMessageListener,
	))
}
