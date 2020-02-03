package controller

import (
	"fmt"
	"github.com/gileslloyd/gs-allocation-service/pkg/rpc"
)

type Allocation struct {
}

func (c Allocation) Execute(message *rpc.Message) (string, error) {
	fmt.Println(message)

	return "", nil
}
