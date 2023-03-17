package alerthook

import (
	"OTRSAlertmanagerHook/logging"
	"OTRSAlertmanagerHook/ticketsystem"
	"encoding/json"
	"net/http"
)

var _ http.Handler = &Manager{}

type Manager struct {
	ticketSystem []ticketsystem.TicketHandler
}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) RegisterTicketSystem(ticketSystem []ticketsystem.TicketHandler) {
	m.ticketSystem = ticketSystem
}

func (m Manager) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	messageDecoder := json.NewDecoder(request.Body)

	var message AlertMassage

	if err := messageDecoder.Decode(&message); err != nil {
		logging.Error(err)
		http.Error(writer, err.Error(), 400)
	}

	m.handleAlert(message)
}

func (m Manager) handleAlert(message AlertMassage) {
	logging.Info(message)
	for _, ticketHandler := range m.ticketSystem {
		for _, alert := range message.Alerts {
			ticketHandler.CreateTicket("["+alert.Status+"] "+alert.Labels["alertname"], alert.Annotations["description"])
		}
	}
}
