package helpers

import "github.com/hmertakyatan/gymapi/ent"

func CalculateProfit(incomes []*ent.Payment, expenses []*ent.Payment) float64 {
	// Gelirleri topla
	totalIncome := 0.0
	for _, income := range incomes {
		totalIncome += income.Amount
	}

	// Giderleri topla
	totalExpense := 0.0
	for _, expense := range expenses {
		totalExpense += expense.Amount
	}

	// Kar-zarar hesapla
	profit := totalIncome - totalExpense

	return profit
}
