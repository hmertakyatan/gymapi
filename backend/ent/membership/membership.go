// Code generated by ent, DO NOT EDIT.

package membership

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the membership type in the database.
	Label = "membership"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOperationDate holds the string denoting the operation_date field in the database.
	FieldOperationDate = "operation_date"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldAmountPaid holds the string denoting the amount_paid field in the database.
	FieldAmountPaid = "amount_paid"
	// FieldAmountRemaining holds the string denoting the amount_remaining field in the database.
	FieldAmountRemaining = "amount_remaining"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// Table holds the table name of the membership in the database.
	Table = "memberships"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "memberships"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_membership"
)

// Columns holds all SQL columns for membership fields.
var Columns = []string{
	FieldID,
	FieldOperationDate,
	FieldStartDate,
	FieldEndDate,
	FieldAmountPaid,
	FieldAmountRemaining,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "memberships"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"customer_membership",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultOperationDate holds the default value on creation for the "operation_date" field.
	DefaultOperationDate time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Membership queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOperationDate orders the results by the operation_date field.
func ByOperationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperationDate, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByAmountPaid orders the results by the amount_paid field.
func ByAmountPaid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountPaid, opts...).ToFunc()
}

// ByAmountRemaining orders the results by the amount_remaining field.
func ByAmountRemaining(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountRemaining, opts...).ToFunc()
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}