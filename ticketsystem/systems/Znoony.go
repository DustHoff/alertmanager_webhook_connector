package systems

import (
	"OTRSAlertmanagerHook/ticketsystem"
)

var _ ticketsystem.TicketHandler = &TicketSystem{}

type TicketSystem struct {
}

func (t TicketSystem) CreateTicket(subject string, body string) {
	//TODO implement me
	panic("implement me")
}

func (t TicketSystem) Init(config ticketsystem.Config) {
}
