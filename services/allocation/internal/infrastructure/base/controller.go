package base

import (
	"github.com/gileslloyd/gs-allocation-service/pkg/rpc"
)

type Controller interface {
	Execute(message *rpc.Message) (string, error)
}
