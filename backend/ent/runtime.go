// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/customer"
	"github.com/hmertakyatan/gymapi/ent/membership"
	"github.com/hmertakyatan/gymapi/ent/membership_type"
	"github.com/hmertakyatan/gymapi/ent/payment"
	"github.com/hmertakyatan/gymapi/ent/personnel"
	"github.com/hmertakyatan/gymapi/ent/personnelrelcustomer"
	"github.com/hmertakyatan/gymapi/ent/schema"
	"github.com/hmertakyatan/gymapi/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescRegistrationDate is the schema descriptor for registration_date field.
	customerDescRegistrationDate := customerFields[3].Descriptor()
	// customer.DefaultRegistrationDate holds the default value on creation for the registration_date field.
	customer.DefaultRegistrationDate = customerDescRegistrationDate.Default.(time.Time)
	// customerDescID is the schema descriptor for id field.
	customerDescID := customerFields[0].Descriptor()
	// customer.DefaultID holds the default value on creation for the id field.
	customer.DefaultID = customerDescID.Default.(func() uuid.UUID)
	membershipFields := schema.Membership{}.Fields()
	_ = membershipFields
	// membershipDescOperationDate is the schema descriptor for operation_date field.
	membershipDescOperationDate := membershipFields[1].Descriptor()
	// membership.DefaultOperationDate holds the default value on creation for the operation_date field.
	membership.DefaultOperationDate = membershipDescOperationDate.Default.(time.Time)
	// membershipDescID is the schema descriptor for id field.
	membershipDescID := membershipFields[0].Descriptor()
	// membership.DefaultID holds the default value on creation for the id field.
	membership.DefaultID = membershipDescID.Default.(func() uuid.UUID)
	membership_typeFields := schema.Membership_type{}.Fields()
	_ = membership_typeFields
	// membership_typeDescID is the schema descriptor for id field.
	membership_typeDescID := membership_typeFields[0].Descriptor()
	// membership_type.DefaultID holds the default value on creation for the id field.
	membership_type.DefaultID = membership_typeDescID.Default.(func() uuid.UUID)
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescDateTime is the schema descriptor for date_time field.
	paymentDescDateTime := paymentFields[4].Descriptor()
	// payment.DefaultDateTime holds the default value on creation for the date_time field.
	payment.DefaultDateTime = paymentDescDateTime.Default.(time.Time)
	// paymentDescID is the schema descriptor for id field.
	paymentDescID := paymentFields[0].Descriptor()
	// payment.DefaultID holds the default value on creation for the id field.
	payment.DefaultID = paymentDescID.Default.(func() uuid.UUID)
	personnelFields := schema.Personnel{}.Fields()
	_ = personnelFields
	// personnelDescStartDate is the schema descriptor for start_date field.
	personnelDescStartDate := personnelFields[5].Descriptor()
	// personnel.DefaultStartDate holds the default value on creation for the start_date field.
	personnel.DefaultStartDate = personnelDescStartDate.Default.(time.Time)
	// personnelDescStatus is the schema descriptor for status field.
	personnelDescStatus := personnelFields[6].Descriptor()
	// personnel.DefaultStatus holds the default value on creation for the status field.
	personnel.DefaultStatus = personnelDescStatus.Default.(bool)
	// personnelDescID is the schema descriptor for id field.
	personnelDescID := personnelFields[0].Descriptor()
	// personnel.DefaultID holds the default value on creation for the id field.
	personnel.DefaultID = personnelDescID.Default.(func() uuid.UUID)
	personnelrelcustomerFields := schema.PersonnelRelCustomer{}.Fields()
	_ = personnelrelcustomerFields
	// personnelrelcustomerDescID is the schema descriptor for id field.
	personnelrelcustomerDescID := personnelrelcustomerFields[0].Descriptor()
	// personnelrelcustomer.DefaultID holds the default value on creation for the id field.
	personnelrelcustomer.DefaultID = personnelrelcustomerDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[6].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
