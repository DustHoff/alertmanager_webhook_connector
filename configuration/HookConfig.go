package configuration

import "OTRSAlertmanagerHook/ticketsystem"

type HookConfig struct {
	TicketSystem []ticketsystem.Config `yaml:"ticketSystem"`
	Address      string                `yaml:"address"`
	Port         string                `yaml:"port"`
	Template     Template              `yaml:"template"`
}

type Template struct {
	Subject  string `yaml:"subject"`
	BodyFile string `yaml:"bodyFile"`
}
