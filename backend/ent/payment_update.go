// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hmertakyatan/gymapi/ent/payment"
	"github.com/hmertakyatan/gymapi/ent/predicate"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetType sets the "type" field.
func (pu *PaymentUpdate) SetType(pa payment.Type) *PaymentUpdate {
	pu.mutation.SetType(pa)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableType(pa *payment.Type) *PaymentUpdate {
	if pa != nil {
		pu.SetType(*pa)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PaymentUpdate) SetDescription(s string) *PaymentUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDescription(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableAmount(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PaymentUpdate) AddAmount(f float64) *PaymentUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetDateTime sets the "date_time" field.
func (pu *PaymentUpdate) SetDateTime(t time.Time) *PaymentUpdate {
	pu.mutation.SetDateTime(t)
	return pu
}

// SetNillableDateTime sets the "date_time" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDateTime(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetDateTime(*t)
	}
	return pu
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Payment.type": %w`, err)}
		}
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(payment.FieldType, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.DateTime(); ok {
		_spec.SetField(payment.FieldDateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetType sets the "type" field.
func (puo *PaymentUpdateOne) SetType(pa payment.Type) *PaymentUpdateOne {
	puo.mutation.SetType(pa)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableType(pa *payment.Type) *PaymentUpdateOne {
	if pa != nil {
		puo.SetType(*pa)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PaymentUpdateOne) SetDescription(s string) *PaymentUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDescription(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableAmount(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetDateTime sets the "date_time" field.
func (puo *PaymentUpdateOne) SetDateTime(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetDateTime(t)
	return puo
}

// SetNillableDateTime sets the "date_time" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDateTime(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetDateTime(*t)
	}
	return puo
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Payment.type": %w`, err)}
		}
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(payment.FieldType, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.DateTime(); ok {
		_spec.SetField(payment.FieldDateTime, field.TypeTime, value)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
