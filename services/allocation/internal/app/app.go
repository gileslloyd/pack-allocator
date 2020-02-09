package app

import (
	"github.com/gileslloyd/gs-allocation-service/config"
)

type App struct {
}

func NewApp() App {
	return App{}
}

func (a App) Run() error {
	messageListener := config.CreateMessageListener()

	messageListener.Listen()

	return nil
}
