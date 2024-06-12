package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	LogLevel string
)

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.protomagic.yaml or ./configs/.protomagic.yaml or .protomagic.yaml)")
	rootCmd.PersistentFlags().StringVarP(&LogLevel, "logLevel", "l", "info", "log level (debug, info, warn, error, fatal, panic)")

	viper.BindPFlag("logLevel", rootCmd.PersistentFlags().Lookup("logLevel"))
	viper.SetDefault("logLevel", "info")
	viper.SetDefault("author", "Carlos Freitas <carlosrfjunior@gmail.com>")
	viper.SetDefault("license", "Apache 2.0")

	rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(debugCmd)

	log.New()

	log.Println(viper.GetString("logLevel"))

	logLevel, err := log.ParseLevel(LogLevel)
	if err != nil {
		fmt.Println("Invalid log level specified:", err)
		os.Exit(1)
	}

	log.Println(logLevel)

	log.SetLevel(logLevel)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(false)
	log.SetOutput(os.Stdout)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	log.Println(viper.GetString("logLevel"))

	Execute()

	log.Println(viper.GetString("logLevel"))
}
