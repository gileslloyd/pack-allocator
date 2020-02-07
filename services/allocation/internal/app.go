package internal

import (
	"fmt"
	"github.com/gileslloyd/gs-allocation-service/internal/infrastructure"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"

	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

type App struct {
	router infrastructure.Router
}

func NewApp(r infrastructure.Router) App {
	return App{
		router: r,
	}
}

func (a App) Start() {
	forever := make(chan struct{})

	go a.listen()

	<-forever
}

func (a App) listen() {
	cmd.Init()

	if err := broker.Init(); err != nil {
		panic(fmt.Sprint("Broker Init error: %v", err))
	}
	if err := broker.Connect(); err != nil {
		panic(fmt.Sprintf("Broker Connect error: %v", err))
	}

	_, err := broker.Subscribe("", func(p broker.Event) error {
		response, err := a.router.Process(string(p.Message().Body))

		if err != nil {
			fmt.Sprintf("ERROR: %v", err)
		}

		broker.Publish(
			"",
			broker.Message{
				Body: []byte(response),
			},
			p.
		)

		return nil
	}, broker.Queue("gs-order"))

	if err != nil {
		panic(err)
	}
}

func (a App) reply() {

}
