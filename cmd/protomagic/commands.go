package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/toolsascode/protomagic/internal/database/mysql"
	"github.com/toolsascode/protomagic/internal/database/postgresql"

	"github.com/spf13/cobra"
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
