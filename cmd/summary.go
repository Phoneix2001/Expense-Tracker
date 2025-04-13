package cmd

import (
	
	"main/dao"

	"github.com/spf13/cobra"
)

var month int32 = 0
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Give Total Expense Amount",
	Run: func(cmd *cobra.Command, args []string) {
		dao.Summary(month)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().Int32VarP(&month, "month", "m", 0, "Give Month Vias Total Amount")
}
