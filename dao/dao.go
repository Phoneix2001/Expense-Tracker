package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"main/model"
	"main/utils/constants"
	"os"
	"slices"
	"text/tabwriter"
	"time"
)

var allExpenses = &[]model.ExpenseTracker{}

func Init(allExpensesParm *[]model.ExpenseTracker) {
	data, err := os.ReadFile(constants.FilePath)
	if err != nil {
		fmt.Println("File not found at the specified path. Creating a new file to store expense tracker data...")
		os.WriteFile(constants.FilePath, make([]byte, 0), 0644)
	}
	// allExpenses = &[]model.ExpenseTracker{}
	errs := json.Unmarshal(data, allExpensesParm)
	if errs != nil {
		log.Print(err)
	}

}

type AddParms struct {
	Desc   string
	Amount int
}

func Add(addParms AddParms) {
	Init(allExpenses)
	Id := 1
	if len(*allExpenses) >= 1 {
		Id = (*allExpenses)[len(*allExpenses)-1].Id + 1
	}
	item := model.ExpenseTracker{
		Id:          Id,
		Date:        time.Now().Format(time.DateTime),
		Description: addParms.Desc,
		Amount:      addParms.Amount,
	}
	*allExpenses = append(*allExpenses, item)

	saveDataToFile()

	fmt.Printf("✅ Expense Created successfully! [ID: %d, Description: %s]\n", item.Id, item.Description)
}

func saveDataToFile() {
	// Step 2: Marshal updated slice
	updatedData, err := json.Marshal(allExpenses)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Step 3: Write it back to the file
	err = os.WriteFile(constants.FilePath, updatedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func List() {
	Init(allExpenses)
	if len(*allExpenses) < 1 {
		fmt.Print("No Record Exist")

	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(writer, "#\tID\tDate\tDescription\tAmount")

	// Rows
	for _, item := range *allExpenses {
		fmt.Fprintf(writer, "#\t%d\t%s\t%s\t%.2f\n", item.Id, item.Date, item.Description, float64(item.Amount))
	}

	// Flush to ensure output is written
	writer.Flush()
}

func Summary(month int32) {
	Init(allExpenses)
	if len(*allExpenses) < 1 {
		fmt.Print("No Record Exist")

	}
	totalAmount := 0
	for _, item := range *allExpenses {
		if month > 0 {
			if expenseCreatedDate, timeErr := time.Parse(time.DateTime, item.Date); timeErr != nil {
				fmt.Print("Some Error Occured")
				break
			} else {
				if int32(expenseCreatedDate.Month()) == month {
					totalAmount += item.Amount
				}
			}
			continue
		}
		totalAmount += item.Amount
	}
	fmt.Printf("# Total expenses: $%d\n", totalAmount)
}

func Delete(id int) {
	Init(allExpenses)
	if len(*allExpenses) < 1 {
		fmt.Print("No Record Exist")

	}
	for index, item := range *allExpenses {
		if item.Id == id {

			*allExpenses = slices.Delete((*allExpenses), index, index+1)
			saveDataToFile()
			fmt.Printf("# Expense deleted successfully")
			break
		}
	}
}
