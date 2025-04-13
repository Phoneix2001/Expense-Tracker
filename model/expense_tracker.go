package model

type ExpenseTracker struct {
	Id          int    `json:"id,omitempty"`
	Date        string `json:"date,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      int    `json:"amount,omitempty"`
}