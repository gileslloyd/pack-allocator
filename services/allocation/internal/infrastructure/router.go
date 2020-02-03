package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure/base"
	"github.com/gileslloyd/gs-allocation-service/pkg/rpc"
)

type Router struct {
	routes map[string]base.Controller
}

func NewRouter(routes map[string]base.Controller) Router {
	return Router{
		routes: routes,
	}
}

func (h Router) Process(message string) (string, error) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(message), &dat); err != nil {
		panic(err)
	}

	controller, err := h.getRoute(dat)

	if err != nil {
		return "", err
	}

	return controller.Execute(rpc.NewMessage(dat["payload"].(map[string]interface{})))
}

func (h Router) getRoute(payload map[string]interface{}) (base.Controller, error) {
	controller := h.routes[fmt.Sprintf("%s.%s", payload["role"].(string), payload["cmd"].(string))]

	if controller != nil {
		return controller, nil
	}

	return nil, errors.New("route not prepared")
}
