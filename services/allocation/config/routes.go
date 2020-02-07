package config

import (
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure/base"
)

var Routes map[string]base.Controller = map[string]base.Controller{
	"pack.allocate": CreateAllocationController(),
}

func GetRoutes() map[string]base.Controller {
	return Routes
}
