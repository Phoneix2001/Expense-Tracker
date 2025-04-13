package cmd

import (
	"fmt"
	"main/dao"

	"github.com/spf13/cobra"
)

var description string
var amount float64

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		if description == "" || amount <= 0 {
			fmt.Println("Both --description and --amount are required and must be valid.")
			return
		}
	  parms :=	dao.AddParms{
		Desc:description,
		Amount: int(amount),
	  }
		dao.Add(parms)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the expense")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Amount spent")
}
