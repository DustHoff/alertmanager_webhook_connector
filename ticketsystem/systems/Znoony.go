package systems

import (
	"OTRSAlertmanagerHook/ticketsystem"
	"OTRSAlertmanagerHook/ticketsystem/systems/znoony"
)

type TicketSystem struct {
	ticketsystem.TicketHandler
	client znoony.Znoony
	config *ticketsystem.Config
}

func (t *TicketSystem) CreateTicket(subject string, body string) (bool, error) {
	_, err := t.client.CreateTicket(znoony.NewZnoonyTicket(t.config.Properties["queue"], subject, body))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *TicketSystem) Init(config *ticketsystem.Config) {
	t.client = znoony.NewZnoonyClient(config)
	t.config = config
}
