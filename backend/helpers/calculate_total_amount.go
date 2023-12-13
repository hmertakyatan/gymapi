package helpers

import "github.com/hmertakyatan/gymapi/ent"

func CalculateTotalAmount(payments []*ent.Payment) float64 {
	totalAmount := 0.0
	for _, payment := range payments {
		totalAmount += payment.Amount
	}
	return totalAmount
}
