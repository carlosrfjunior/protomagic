package main

import (
	"fmt"

	"github.com/carlosrfjunior/protomagic/internal/database/mysql"
	"github.com/carlosrfjunior/protomagic/internal/database/postgresql"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "protomagic",
	Short: "ProtoMagic is a CLI that helps convert database tables into Protocol Buffers files.",
	Long: `ProtoMagic is a CLI that helps convert database tables into Protocol Buffers files.
	Complete documentation is available at https://protomagic.dev`,
	Run: func(cmd *cobra.Command, args []string) {
		postgresql.Generate().Run()
		mysql.Generate().Run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ProtoMagic",
	Long:  `All software has versions. This is ProtoMagic's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:\t", Version)
		fmt.Println("build.Time:\t", Time)
		fmt.Println("build.User:\t", User)
	},
}
