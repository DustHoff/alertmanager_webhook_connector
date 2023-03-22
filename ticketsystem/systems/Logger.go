package systems

import (
	"OTRSAlertmanagerHook/logging"
	"OTRSAlertmanagerHook/ticketsystem"
	"log"
)

type LoggerSystem struct {
	ticketsystem.TicketHandler
	logger log.Logger
}

func (l LoggerSystem) Init(config *ticketsystem.Config) {
}

func (l LoggerSystem) CreateTicket(subject string, body string) (bool, error) {
	logging.Info("Subject", subject)
	logging.Info("Body", body)
	return true, nil
}
