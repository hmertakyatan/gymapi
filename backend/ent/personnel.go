// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/personnel"
)

// Personnel is the model entity for the Personnel schema.
type Personnel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Salary holds the value of the "salary" field.
	Salary float64 `json:"salary,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// BithDate holds the value of the "bith_date" field.
	BithDate time.Time `json:"bith_date,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonnelQuery when eager-loading is set.
	Edges        PersonnelEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PersonnelEdges holds the relations/edges for other nodes in the graph.
type PersonnelEdges struct {
	// PersonnelRelCustomer holds the value of the personnel_rel_customer edge.
	PersonnelRelCustomer []*PersonnelRelCustomer `json:"personnel_rel_customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PersonnelRelCustomerOrErr returns the PersonnelRelCustomer value or an error if the edge
// was not loaded in eager-loading.
func (e PersonnelEdges) PersonnelRelCustomerOrErr() ([]*PersonnelRelCustomer, error) {
	if e.loadedTypes[0] {
		return e.PersonnelRelCustomer, nil
	}
	return nil, &NotLoadedError{edge: "personnel_rel_customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Personnel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case personnel.FieldStatus:
			values[i] = new(sql.NullBool)
		case personnel.FieldSalary:
			values[i] = new(sql.NullFloat64)
		case personnel.FieldName, personnel.FieldDescription:
			values[i] = new(sql.NullString)
		case personnel.FieldBithDate, personnel.FieldStartDate:
			values[i] = new(sql.NullTime)
		case personnel.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Personnel fields.
func (pe *Personnel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personnel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pe.ID = *value
			}
		case personnel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case personnel.FieldSalary:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field salary", values[i])
			} else if value.Valid {
				pe.Salary = value.Float64
			}
		case personnel.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case personnel.FieldBithDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field bith_date", values[i])
			} else if value.Valid {
				pe.BithDate = value.Time
			}
		case personnel.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				pe.StartDate = value.Time
			}
		case personnel.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pe.Status = value.Bool
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Personnel.
// This includes values selected through modifiers, order, etc.
func (pe *Personnel) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryPersonnelRelCustomer queries the "personnel_rel_customer" edge of the Personnel entity.
func (pe *Personnel) QueryPersonnelRelCustomer() *PersonnelRelCustomerQuery {
	return NewPersonnelClient(pe.config).QueryPersonnelRelCustomer(pe)
}

// Update returns a builder for updating this Personnel.
// Note that you need to call Personnel.Unwrap() before calling this method if this Personnel
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Personnel) Update() *PersonnelUpdateOne {
	return NewPersonnelClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Personnel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Personnel) Unwrap() *Personnel {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Personnel is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Personnel) String() string {
	var builder strings.Builder
	builder.WriteString("Personnel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("salary=")
	builder.WriteString(fmt.Sprintf("%v", pe.Salary))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("bith_date=")
	builder.WriteString(pe.BithDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(pe.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pe.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Personnels is a parsable slice of Personnel.
type Personnels []*Personnel
