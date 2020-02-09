package base

import (
	"github.com/gileslloyd/gs-allocation-service/pkg/infrastructure"
)

type Controller interface {
	Execute(message *infrastructure.Message) (string, error)
}
