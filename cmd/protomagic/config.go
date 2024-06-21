package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func initConfig() {

	log.New()

	logLevel, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		fmt.Println("Invalid log level specified:", err)
		os.Exit(1)
	}

	log.SetLevel(logLevel)

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(".")
		viper.SetConfigName(".protomagic")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Can't read config:", err)
		os.Exit(1)
	}
}
