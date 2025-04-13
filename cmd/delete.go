package cmd

import (
	"fmt"
	"main/dao"

	"github.com/spf13/cobra"
)


var Id float64

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a expense",
	Run: func(cmd *cobra.Command, args []string) {
		if  Id < 0 {
			fmt.Println("--id are required and must be valid.")
			return
		}
	 
		dao.Delete(int(Id))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Float64VarP(&Id, "id", "i", 0.0, "Expense Id")
}
