package configuration

import (
	"OTRSAlertmanagerHook/logging"
	"gopkg.in/yaml.v3"
	"os"
)

func Load(path *string) HookConfig {
	var configFile = *path
	logging.Info("loading ", configFile, " as configuration")

	file, err := os.ReadFile(configFile)
	if err != nil {
		logging.Fatal(err)
	}
	var config HookConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		logging.Fatal(err)
	}
	return config
}
