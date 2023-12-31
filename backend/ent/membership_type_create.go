// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/membership_type"
)

// MembershipTypeCreate is the builder for creating a Membership_type entity.
type MembershipTypeCreate struct {
	config
	mutation *MembershipTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mtc *MembershipTypeCreate) SetName(s string) *MembershipTypeCreate {
	mtc.mutation.SetName(s)
	return mtc
}

// SetDescription sets the "description" field.
func (mtc *MembershipTypeCreate) SetDescription(s string) *MembershipTypeCreate {
	mtc.mutation.SetDescription(s)
	return mtc
}

// SetMembershipMonth sets the "membership_month" field.
func (mtc *MembershipTypeCreate) SetMembershipMonth(u uint8) *MembershipTypeCreate {
	mtc.mutation.SetMembershipMonth(u)
	return mtc
}

// SetPrice sets the "price" field.
func (mtc *MembershipTypeCreate) SetPrice(f float64) *MembershipTypeCreate {
	mtc.mutation.SetPrice(f)
	return mtc
}

// SetID sets the "id" field.
func (mtc *MembershipTypeCreate) SetID(u uuid.UUID) *MembershipTypeCreate {
	mtc.mutation.SetID(u)
	return mtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mtc *MembershipTypeCreate) SetNillableID(u *uuid.UUID) *MembershipTypeCreate {
	if u != nil {
		mtc.SetID(*u)
	}
	return mtc
}

// Mutation returns the MembershipTypeMutation object of the builder.
func (mtc *MembershipTypeCreate) Mutation() *MembershipTypeMutation {
	return mtc.mutation
}

// Save creates the Membership_type in the database.
func (mtc *MembershipTypeCreate) Save(ctx context.Context) (*Membership_type, error) {
	mtc.defaults()
	return withHooks(ctx, mtc.sqlSave, mtc.mutation, mtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MembershipTypeCreate) SaveX(ctx context.Context) *Membership_type {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtc *MembershipTypeCreate) Exec(ctx context.Context) error {
	_, err := mtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtc *MembershipTypeCreate) ExecX(ctx context.Context) {
	if err := mtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mtc *MembershipTypeCreate) defaults() {
	if _, ok := mtc.mutation.ID(); !ok {
		v := membership_type.DefaultID()
		mtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MembershipTypeCreate) check() error {
	if _, ok := mtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Membership_type.name"`)}
	}
	if _, ok := mtc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Membership_type.description"`)}
	}
	if _, ok := mtc.mutation.MembershipMonth(); !ok {
		return &ValidationError{Name: "membership_month", err: errors.New(`ent: missing required field "Membership_type.membership_month"`)}
	}
	if _, ok := mtc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Membership_type.price"`)}
	}
	return nil
}

func (mtc *MembershipTypeCreate) sqlSave(ctx context.Context) (*Membership_type, error) {
	if err := mtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
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
	mtc.mutation.id = &_node.ID
	mtc.mutation.done = true
	return _node, nil
}

func (mtc *MembershipTypeCreate) createSpec() (*Membership_type, *sqlgraph.CreateSpec) {
	var (
		_node = &Membership_type{config: mtc.config}
		_spec = sqlgraph.NewCreateSpec(membership_type.Table, sqlgraph.NewFieldSpec(membership_type.FieldID, field.TypeUUID))
	)
	if id, ok := mtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mtc.mutation.Name(); ok {
		_spec.SetField(membership_type.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mtc.mutation.Description(); ok {
		_spec.SetField(membership_type.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mtc.mutation.MembershipMonth(); ok {
		_spec.SetField(membership_type.FieldMembershipMonth, field.TypeUint8, value)
		_node.MembershipMonth = value
	}
	if value, ok := mtc.mutation.Price(); ok {
		_spec.SetField(membership_type.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	return _node, _spec
}

// MembershipTypeCreateBulk is the builder for creating many Membership_type entities in bulk.
type MembershipTypeCreateBulk struct {
	config
	err      error
	builders []*MembershipTypeCreate
}

// Save creates the Membership_type entities in the database.
func (mtcb *MembershipTypeCreateBulk) Save(ctx context.Context) ([]*Membership_type, error) {
	if mtcb.err != nil {
		return nil, mtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*Membership_type, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MembershipTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MembershipTypeCreateBulk) SaveX(ctx context.Context) []*Membership_type {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtcb *MembershipTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := mtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtcb *MembershipTypeCreateBulk) ExecX(ctx context.Context) {
	if err := mtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
