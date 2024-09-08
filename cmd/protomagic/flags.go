package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile            string
	logLevel           string
	protoPath          string
	protoPathReset     bool
	protoTemplateFile  string
	protoSyntaxVersion string
	ApiVersion         string
)

func initFlag() {

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.protomagic.yaml or ./configs/.protomagic.yaml or .protomagic.yaml)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "log level (debug, info, warn, error, fatal, panic)")
	rootCmd.PersistentFlags().StringVarP(&protoPath, "proto-path-output", "o", "proto", "Path for proto files.")
	rootCmd.PersistentFlags().StringVarP(&protoTemplateFile, "proto-template-file", "t", "", "Example: proto_template.gotmpl")
	rootCmd.PersistentFlags().StringVarP(&protoSyntaxVersion, "proto-syntax-version", "V", "proto3", "The version of the protocol buffers in your project.")
	rootCmd.PersistentFlags().StringVarP(&ApiVersion, "proto-api-version", "v", "v1", "Your project's API version.")
	rootCmd.PersistentFlags().BoolVarP(&protoPathReset, "proto-path-reset", "R", false, "Reset the path for proto files.")

	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("protobuf.output.path", rootCmd.PersistentFlags().Lookup("proto-path-output"))
	viper.BindPFlag("protobuf.output.reset", rootCmd.PersistentFlags().Lookup("proto-path-reset"))
	viper.BindPFlag("protobuf.templateFile", rootCmd.PersistentFlags().Lookup("proto-template-file"))
	viper.BindPFlag("protobuf.syntax", rootCmd.PersistentFlags().Lookup("proto-syntax-version"))
	viper.BindPFlag("protobuf.apiVersion", rootCmd.PersistentFlags().Lookup("proto-api-version"))

	viper.SetDefault("author", "Carlos Freitas <carlosrfjunior@gmail.com>")
	viper.SetDefault("license", "Apache 2.0")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(docsCmd)

}
