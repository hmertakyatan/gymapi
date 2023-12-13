// Code generated by ent, DO NOT EDIT.

package customer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldFullName, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldBirthDate, v))
}

// RegistrationDate applies equality check predicate on the "registration_date" field. It's identical to RegistrationDateEQ.
func RegistrationDate(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldRegistrationDate, v))
}

// CustomerPictureURL applies equality check predicate on the "customer_picture_url" field. It's identical to CustomerPictureURLEQ.
func CustomerPictureURL(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCustomerPictureURL, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldStatus, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldFullName, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldBirthDate, v))
}

// RegistrationDateEQ applies the EQ predicate on the "registration_date" field.
func RegistrationDateEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldRegistrationDate, v))
}

// RegistrationDateNEQ applies the NEQ predicate on the "registration_date" field.
func RegistrationDateNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldRegistrationDate, v))
}

// RegistrationDateIn applies the In predicate on the "registration_date" field.
func RegistrationDateIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldRegistrationDate, vs...))
}

// RegistrationDateNotIn applies the NotIn predicate on the "registration_date" field.
func RegistrationDateNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldRegistrationDate, vs...))
}

// RegistrationDateGT applies the GT predicate on the "registration_date" field.
func RegistrationDateGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldRegistrationDate, v))
}

// RegistrationDateGTE applies the GTE predicate on the "registration_date" field.
func RegistrationDateGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldRegistrationDate, v))
}

// RegistrationDateLT applies the LT predicate on the "registration_date" field.
func RegistrationDateLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldRegistrationDate, v))
}

// RegistrationDateLTE applies the LTE predicate on the "registration_date" field.
func RegistrationDateLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldRegistrationDate, v))
}

// CustomerPictureURLEQ applies the EQ predicate on the "customer_picture_url" field.
func CustomerPictureURLEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCustomerPictureURL, v))
}

// CustomerPictureURLNEQ applies the NEQ predicate on the "customer_picture_url" field.
func CustomerPictureURLNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCustomerPictureURL, v))
}

// CustomerPictureURLIn applies the In predicate on the "customer_picture_url" field.
func CustomerPictureURLIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCustomerPictureURL, vs...))
}

// CustomerPictureURLNotIn applies the NotIn predicate on the "customer_picture_url" field.
func CustomerPictureURLNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCustomerPictureURL, vs...))
}

// CustomerPictureURLGT applies the GT predicate on the "customer_picture_url" field.
func CustomerPictureURLGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCustomerPictureURL, v))
}

// CustomerPictureURLGTE applies the GTE predicate on the "customer_picture_url" field.
func CustomerPictureURLGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCustomerPictureURL, v))
}

// CustomerPictureURLLT applies the LT predicate on the "customer_picture_url" field.
func CustomerPictureURLLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCustomerPictureURL, v))
}

// CustomerPictureURLLTE applies the LTE predicate on the "customer_picture_url" field.
func CustomerPictureURLLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCustomerPictureURL, v))
}

// CustomerPictureURLContains applies the Contains predicate on the "customer_picture_url" field.
func CustomerPictureURLContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldCustomerPictureURL, v))
}

// CustomerPictureURLHasPrefix applies the HasPrefix predicate on the "customer_picture_url" field.
func CustomerPictureURLHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldCustomerPictureURL, v))
}

// CustomerPictureURLHasSuffix applies the HasSuffix predicate on the "customer_picture_url" field.
func CustomerPictureURLHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldCustomerPictureURL, v))
}

// CustomerPictureURLIsNil applies the IsNil predicate on the "customer_picture_url" field.
func CustomerPictureURLIsNil() predicate.Customer {
	return predicate.Customer(sql.FieldIsNull(FieldCustomerPictureURL))
}

// CustomerPictureURLNotNil applies the NotNil predicate on the "customer_picture_url" field.
func CustomerPictureURLNotNil() predicate.Customer {
	return predicate.Customer(sql.FieldNotNull(FieldCustomerPictureURL))
}

// CustomerPictureURLEqualFold applies the EqualFold predicate on the "customer_picture_url" field.
func CustomerPictureURLEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldCustomerPictureURL, v))
}

// CustomerPictureURLContainsFold applies the ContainsFold predicate on the "customer_picture_url" field.
func CustomerPictureURLContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldCustomerPictureURL, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldStatus, v))
}

// HasMembership applies the HasEdge predicate on the "membership" edge.
func HasMembership() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MembershipTable, MembershipColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembershipWith applies the HasEdge predicate on the "membership" edge with a given conditions (other predicates).
func HasMembershipWith(preds ...predicate.Membership) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newMembershipStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonnelRelCustomer applies the HasEdge predicate on the "personnel_rel_customer" edge.
func HasPersonnelRelCustomer() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonnelRelCustomerTable, PersonnelRelCustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonnelRelCustomerWith applies the HasEdge predicate on the "personnel_rel_customer" edge with a given conditions (other predicates).
func HasPersonnelRelCustomerWith(preds ...predicate.PersonnelRelCustomer) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newPersonnelRelCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.NotPredicates(p))
}
