package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var DEFAULT_CFG_FILE_LOOKUP_NAME = ".default_conf"

func InitDefaultViperConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(DEFAULT_CFG_FILE_LOOKUP_NAME)
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
