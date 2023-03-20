package systems

import (
	"OTRSAlertmanagerHook/ticketsystem"
	"log"
)

type LoggerSystem struct {
	ticketsystem.TicketHandler
	logger log.Logger
}

func (l LoggerSystem) Init(config ticketsystem.Config) {
}

func (l LoggerSystem) CreateTicket(subject string, body string) {
	log.Println("test")
}
