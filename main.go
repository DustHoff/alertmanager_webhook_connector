package main

import (
	"OTRSAlertmanagerHook/alerthook"
	"OTRSAlertmanagerHook/configuration"
	"OTRSAlertmanagerHook/logging"
	"OTRSAlertmanagerHook/ticketsystem/manager"
	"flag"
	"net/http"
)

func main() {
	configFile := flag.String("config", "config.yml", "")
	flag.Parse()
	config := configuration.Load(configFile)

	hook := alerthook.NewManager()
	hook.RegisterTicketSystem(manager.CreateTicketHandler(config.TicketSystem))
	http.Handle("/alert", hook)
	logging.Error(http.ListenAndServe(":8080", nil))
}
