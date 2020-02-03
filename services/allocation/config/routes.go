package config

import (
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure/base"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure/controller"
)

var Routes map[string]base.Controller = map[string]base.Controller{
	"pack.allocate": controller.Allocation{},
}

func GetRoutes() map[string]base.Controller {
	return Routes
}
