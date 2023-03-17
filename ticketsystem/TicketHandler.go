package ticketsystem

type TicketHandler interface {
	Init(config Config)
	CreateTicket(subject string, body string)
}
