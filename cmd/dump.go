package cmd

import (
	"github.com/spf13/viper"
	"github.com/openlab-red/etcd-data-migration/pkg/dump"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump etcd data in a file",
	Long:  `Dump etcd data in a file`,
	Run: func(cmd *cobra.Command, args []string) {
		dump.Start()
	},
}

func init() {
	RootCmd.AddCommand(dumpCmd)

	dumpCmd.Flags().StringP("endpoint","e", "localhost:2379", "Etcd endpoint")
	dumpCmd.Flags().StringP("output","o", "etcd.dump", "Dump file output")
	viper.BindPFlag("endpoint", dumpCmd.Flags().Lookup("endpoint"))
	viper.BindPFlag("output", dumpCmd.Flags().Lookup("output"))
}

