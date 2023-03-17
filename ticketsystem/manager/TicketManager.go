package manager

import (
	"OTRSAlertmanagerHook/logging"
	"OTRSAlertmanagerHook/ticketsystem"
	"OTRSAlertmanagerHook/ticketsystem/systems"
)

func getType(t string) ticketsystem.TicketHandler {
	switch t {
	case "ZNOONY":
		return systems.TicketSystem{}
	case "LOGGER":
		return systems.LoggerSystem{}
	default:
		logging.Fatal("Unknown Type ", t)
	}
	return nil
}

func CreateTicketHandler(configList []ticketsystem.Config) []ticketsystem.TicketHandler {
	var allHandler []ticketsystem.TicketHandler
	for _, config := range configList {
		logging.Info("Initializing Connector", config.Type)
		handler := getType(config.Type)
		handler.Init(config)
		allHandler = append(allHandler, handler)
	}
	return allHandler
}
