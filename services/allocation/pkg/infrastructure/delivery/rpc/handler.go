package rpc

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gileslloyd/gs-allocation-service/internal/app/base"
	"github.com/gileslloyd/gs-allocation-service/pkg/infrastructure"
)

type Handler struct {
	routes map[string]base.Controller
}

func NewHandler(routes map[string]base.Controller) *Handler {
	return &Handler{
		routes: routes,
	}
}

func (h Handler) Process(message string) (string, error) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(message), &dat); err != nil {
		panic(err)
	}

	controller, err := h.getRoute(dat)

	if err != nil {
		return "", err
	}

	return controller.Execute(infrastructure.NewMessage(dat["payload"].(map[string]interface{})))
}

func (h Handler) getRoute(payload map[string]interface{}) (base.Controller, error) {
	controller := h.routes[fmt.Sprintf("%s.%s", payload["role"].(string), payload["cmd"].(string))]

	if controller != nil {
		return controller, nil
	}

	return nil, errors.New("route not prepared")
}
