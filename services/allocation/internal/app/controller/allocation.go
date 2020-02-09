package controller

import (
	"encoding/json"
	"github.com/gileslloyd/gs-allocation-service/internal/domain/allocation"
	"github.com/gileslloyd/gs-allocation-service/pkg/infrastructure"
)

type Allocation struct {
	service allocation.Service
}

func NewAllocationController(service allocation.Service) Allocation {
	return Allocation{service: service}
}

func (c Allocation) Execute(message *infrastructure.Message) (string, error) {
	requiredItems := int(message.Get("requiredItems", "0").(float64))

	response, err := json.Marshal(c.service.GetPackAllocation(requiredItems))

	if err != nil {
		return "", err
	}

	response, err = json.Marshal(
		map[string]string{ "data": string(response) },
	)

	if err != nil {
		return "", err
	}

	return string(response), nil
}
