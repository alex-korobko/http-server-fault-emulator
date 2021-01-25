package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"log"
)

const BEHAVIOUR_KEY = "Behaviour"
const EMULATION_PORT_KEY = "EmulationPort"

func Init() {
	setDefaultValues()
	//load from config file
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".") // look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; try to create a new one with default values
			safeWriteErr := viper.SafeWriteConfig()
			if (safeWriteErr != nil) {
				log.Printf("Safe write config failed: ", safeWriteErr.Error())
			}
		}
		log.Fatal(fmt.Sprintf("Config init error: %s", err.Error()))
	}
	viper.WatchConfig()
	viper.OnConfigChange(reReadConfig)
}

func setDefaultValues() {
	//default values
	viper.SetDefault(EMULATION_PORT_KEY, "7777")
	viper.SetDefault(BEHAVIOUR_KEY, "close_connection")
}

func reReadConfig(e fsnotify.Event) {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


func GetEmulationPort() string {
	return viper.GetString(EMULATION_PORT_KEY)
}

func GetCurrentBehaviourName() string {
	return viper.GetString(BEHAVIOUR_KEY)
}

func GetCurrentBehaviourParams() map[string]interface{} {
	return viper.GetStringMap(GetCurrentBehaviourName())
}



