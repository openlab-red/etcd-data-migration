package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/openlab-red/etcd-data-migration/pkg/clone"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone data from source etcd to target etcd",
	Long:  `Clone data from source etcd to target etcd`,
	Run: func(cmd *cobra.Command, args []string) {
		clone.Start()
	},
}

func init() {
	RootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringP("source","s", "localhost:2379", "Source Etcd endpoint")
	cloneCmd.Flags().StringP("target","t", "localhost:3379", "Target Etcd endpoint")
	cloneCmd.Flags().StringP("overwrite","o", "false", "Overwrite existing keys")
	viper.BindPFlag("source", cloneCmd.Flags().Lookup("source"))
	viper.BindPFlag("target", cloneCmd.Flags().Lookup("target"))
}
