package services

import "github.com/hmertakyatan/gymapi/ent"

type DiscountService interface {
	BirthDateDiscount(c ent.Customer, m ent.Membership)
}
