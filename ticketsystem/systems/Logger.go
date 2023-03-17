package systems

import (
	"OTRSAlertmanagerHook/ticketsystem"
	"log"
	"os"
)

var _ ticketsystem.TicketHandler = &LoggerSystem{}

type LoggerSystem struct {
	logger *log.Logger
}

func (l LoggerSystem) Init(config ticketsystem.Config) {
	l.logger = log.New(os.Stdout, "Alert:", log.Ldate|log.Ltime)
}

func (l LoggerSystem) CreateTicket(subject string, body string) {
	l.logger.Println(subject, body)
}
