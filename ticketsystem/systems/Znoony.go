package systems

import (
	"OTRSAlertmanagerHook/ticketsystem"
	"net/http"
	"time"
)

type TicketSystem struct {
	ticketsystem.TicketHandler
	client http.Client
}

func (t TicketSystem) CreateTicket(subject string, body string) {
	//TODO implement me
	panic("implement me")
}

func (t TicketSystem) Init(config ticketsystem.Config) {
	t.client = http.Client{
		Timeout: 30 * time.Second,
	}
}
