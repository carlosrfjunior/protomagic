package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/toolsascode/protomagic/internal/database/mysql"
	"github.com/toolsascode/protomagic/internal/database/postgresql"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "protomagic",
	Short: "ProtoMagic is a CLI that helps convert database tables into Protocol Buffers files.",
	Long: `ProtoMagic is a CLI that helps convert database tables into Protocol Buffers files.
	Complete documentation is available at https://protomagic.dev`,
	Run: func(cmd *cobra.Command, args []string) {

		path := viper.GetString("protobuf.output.path")

		log.Debugf("The path output for proto files: %s", path)

		if viper.GetBool("protobuf.output.reset") {
			log.Debugf("You have chosen to recreate the directory: %s", path)
			os.RemoveAll(path)
		}

		postgresql.Generate().Run()
		mysql.Generate().Run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ProtoMagic",
	Long:  `All software has versions. This is ProtoMagic's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: ", version)
		fmt.Println("Date: ", date)
		fmt.Println("Commit: ", commit)
		fmt.Println("Built by: ", builtBy)
	},
}

var docsCmd = &cobra.Command{
	Use:    "docs",
	Short:  "Generating ProtoMagic CLI markdown documentation.",
	Long:   `Allow generating documentation in markdown format for ProtoMagic CLI internal commands`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		var path = "./docs"

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Generating markdown documentation")
		err := doc.GenMarkdownTree(rootCmd, path)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Documentation successfully generated in %s", path)
	},
}
