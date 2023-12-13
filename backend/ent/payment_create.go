// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/payment"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (pc *PaymentCreate) SetType(pa payment.Type) *PaymentCreate {
	pc.mutation.SetType(pa)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PaymentCreate) SetDescription(s string) *PaymentCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(f float64) *PaymentCreate {
	pc.mutation.SetAmount(f)
	return pc
}

// SetDateTime sets the "date_time" field.
func (pc *PaymentCreate) SetDateTime(t time.Time) *PaymentCreate {
	pc.mutation.SetDateTime(t)
	return pc
}

// SetNillableDateTime sets the "date_time" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableDateTime(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetDateTime(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PaymentCreate) SetID(u uuid.UUID) *PaymentCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableID(u *uuid.UUID) *PaymentCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PaymentCreate) defaults() {
	if _, ok := pc.mutation.DateTime(); !ok {
		v := payment.DefaultDateTime
		pc.mutation.SetDateTime(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := payment.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Payment.type"`)}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Payment.type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Payment.description"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Payment.amount"`)}
	}
	if _, ok := pc.mutation.DateTime(); !ok {
		return &ValidationError{Name: "date_time", err: errors.New(`ent: missing required field "Payment.date_time"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(payment.Table, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(payment.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.DateTime(); ok {
		_spec.SetField(payment.FieldDateTime, field.TypeTime, value)
		_node.DateTime = value
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	err      error
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
