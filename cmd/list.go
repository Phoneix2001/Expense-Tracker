package cmd

import (
	"main/dao"

	"github.com/spf13/cobra"
)



var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		dao.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
