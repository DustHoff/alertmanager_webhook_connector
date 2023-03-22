package alerthook

import (
	"OTRSAlertmanagerHook/configuration"
	"OTRSAlertmanagerHook/logging"
	"OTRSAlertmanagerHook/ticketsystem"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"text/template"
)

var _ http.Handler = &Manager{}

type Manager struct {
	sync.Mutex
	ticketSystem    *[]ticketsystem.TicketHandler
	subjectTemplate *template.Template
	bodyTemplate    *template.Template
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) RegisterTicketSystem(ticketSystem *[]ticketsystem.TicketHandler) {
	m.Lock()
	defer m.Unlock()
	m.ticketSystem = ticketSystem
}

func (m *Manager) SetTemplate(t configuration.Template) {
	m.Lock()
	defer m.Unlock()
	templateEngine := template.New("subjectEngine")
	var err error
	m.subjectTemplate, err = templateEngine.Parse(t.Subject)
	if err != nil {
		logging.Fatal(err)
	}
	templateEngine = template.New("bodyEngine")
	bodyTemplate, err := os.ReadFile(t.BodyFile)
	if err != nil {
		logging.Fatal(err)
	}
	m.bodyTemplate, err = templateEngine.Parse(string(bodyTemplate))
	if err != nil {
		logging.Fatal(err)
	}
}

func (m *Manager) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func (m *Manager) handleAlert(message AlertMassage) {
	m.Lock()
	defer m.Unlock()

	for _, ticketHandler := range *m.ticketSystem {
		for _, alert := range message.Alerts {
			subjectBuffer := bytes.NewBufferString("")
			bodyBuffer := bytes.NewBufferString("")
			err := m.subjectTemplate.Execute(subjectBuffer, alert)
			if err != nil {
				logging.Error(err)
			}
			err = m.bodyTemplate.Execute(bodyBuffer, alert)
			if err != nil {
				logging.Error(err)
			}
			success, err := ticketHandler.CreateTicket(subjectBuffer.String(), bodyBuffer.String())
			logging.Info("Ticket creation successfully:", success)
			if err != nil {
				logging.Error(err)
			}
		}
	}
}
