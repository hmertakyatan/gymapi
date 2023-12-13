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
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/customer"
	"github.com/hmertakyatan/gymapi/ent/membership"
	"github.com/hmertakyatan/gymapi/ent/predicate"
)

// MembershipUpdate is the builder for updating Membership entities.
type MembershipUpdate struct {
	config
	hooks    []Hook
	mutation *MembershipMutation
}

// Where appends a list predicates to the MembershipUpdate builder.
func (mu *MembershipUpdate) Where(ps ...predicate.Membership) *MembershipUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetOperationDate sets the "operation_date" field.
func (mu *MembershipUpdate) SetOperationDate(t time.Time) *MembershipUpdate {
	mu.mutation.SetOperationDate(t)
	return mu
}

// SetNillableOperationDate sets the "operation_date" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableOperationDate(t *time.Time) *MembershipUpdate {
	if t != nil {
		mu.SetOperationDate(*t)
	}
	return mu
}

// SetStartDate sets the "start_date" field.
func (mu *MembershipUpdate) SetStartDate(t time.Time) *MembershipUpdate {
	mu.mutation.SetStartDate(t)
	return mu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableStartDate(t *time.Time) *MembershipUpdate {
	if t != nil {
		mu.SetStartDate(*t)
	}
	return mu
}

// SetEndDate sets the "end_date" field.
func (mu *MembershipUpdate) SetEndDate(t time.Time) *MembershipUpdate {
	mu.mutation.SetEndDate(t)
	return mu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableEndDate(t *time.Time) *MembershipUpdate {
	if t != nil {
		mu.SetEndDate(*t)
	}
	return mu
}

// SetAmountPaid sets the "amount_paid" field.
func (mu *MembershipUpdate) SetAmountPaid(f float64) *MembershipUpdate {
	mu.mutation.ResetAmountPaid()
	mu.mutation.SetAmountPaid(f)
	return mu
}

// SetNillableAmountPaid sets the "amount_paid" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableAmountPaid(f *float64) *MembershipUpdate {
	if f != nil {
		mu.SetAmountPaid(*f)
	}
	return mu
}

// AddAmountPaid adds f to the "amount_paid" field.
func (mu *MembershipUpdate) AddAmountPaid(f float64) *MembershipUpdate {
	mu.mutation.AddAmountPaid(f)
	return mu
}

// ClearAmountPaid clears the value of the "amount_paid" field.
func (mu *MembershipUpdate) ClearAmountPaid() *MembershipUpdate {
	mu.mutation.ClearAmountPaid()
	return mu
}

// SetAmountRemaining sets the "amount_remaining" field.
func (mu *MembershipUpdate) SetAmountRemaining(f float64) *MembershipUpdate {
	mu.mutation.ResetAmountRemaining()
	mu.mutation.SetAmountRemaining(f)
	return mu
}

// SetNillableAmountRemaining sets the "amount_remaining" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableAmountRemaining(f *float64) *MembershipUpdate {
	if f != nil {
		mu.SetAmountRemaining(*f)
	}
	return mu
}

// AddAmountRemaining adds f to the "amount_remaining" field.
func (mu *MembershipUpdate) AddAmountRemaining(f float64) *MembershipUpdate {
	mu.mutation.AddAmountRemaining(f)
	return mu
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (mu *MembershipUpdate) SetCustomerID(id uuid.UUID) *MembershipUpdate {
	mu.mutation.SetCustomerID(id)
	return mu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (mu *MembershipUpdate) SetCustomer(c *Customer) *MembershipUpdate {
	return mu.SetCustomerID(c.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (mu *MembershipUpdate) Mutation() *MembershipMutation {
	return mu.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (mu *MembershipUpdate) ClearCustomer() *MembershipUpdate {
	mu.mutation.ClearCustomer()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MembershipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MembershipUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MembershipUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MembershipUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MembershipUpdate) check() error {
	if _, ok := mu.mutation.CustomerID(); mu.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Membership.customer"`)
	}
	return nil
}

func (mu *MembershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.OperationDate(); ok {
		_spec.SetField(membership.FieldOperationDate, field.TypeTime, value)
	}
	if value, ok := mu.mutation.StartDate(); ok {
		_spec.SetField(membership.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := mu.mutation.EndDate(); ok {
		_spec.SetField(membership.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := mu.mutation.AmountPaid(); ok {
		_spec.SetField(membership.FieldAmountPaid, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.AddedAmountPaid(); ok {
		_spec.AddField(membership.FieldAmountPaid, field.TypeFloat64, value)
	}
	if mu.mutation.AmountPaidCleared() {
		_spec.ClearField(membership.FieldAmountPaid, field.TypeFloat64)
	}
	if value, ok := mu.mutation.AmountRemaining(); ok {
		_spec.SetField(membership.FieldAmountRemaining, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.AddedAmountRemaining(); ok {
		_spec.AddField(membership.FieldAmountRemaining, field.TypeFloat64, value)
	}
	if mu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.CustomerTable,
			Columns: []string{membership.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.CustomerTable,
			Columns: []string{membership.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MembershipUpdateOne is the builder for updating a single Membership entity.
type MembershipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MembershipMutation
}

// SetOperationDate sets the "operation_date" field.
func (muo *MembershipUpdateOne) SetOperationDate(t time.Time) *MembershipUpdateOne {
	muo.mutation.SetOperationDate(t)
	return muo
}

// SetNillableOperationDate sets the "operation_date" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableOperationDate(t *time.Time) *MembershipUpdateOne {
	if t != nil {
		muo.SetOperationDate(*t)
	}
	return muo
}

// SetStartDate sets the "start_date" field.
func (muo *MembershipUpdateOne) SetStartDate(t time.Time) *MembershipUpdateOne {
	muo.mutation.SetStartDate(t)
	return muo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableStartDate(t *time.Time) *MembershipUpdateOne {
	if t != nil {
		muo.SetStartDate(*t)
	}
	return muo
}

// SetEndDate sets the "end_date" field.
func (muo *MembershipUpdateOne) SetEndDate(t time.Time) *MembershipUpdateOne {
	muo.mutation.SetEndDate(t)
	return muo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableEndDate(t *time.Time) *MembershipUpdateOne {
	if t != nil {
		muo.SetEndDate(*t)
	}
	return muo
}

// SetAmountPaid sets the "amount_paid" field.
func (muo *MembershipUpdateOne) SetAmountPaid(f float64) *MembershipUpdateOne {
	muo.mutation.ResetAmountPaid()
	muo.mutation.SetAmountPaid(f)
	return muo
}

// SetNillableAmountPaid sets the "amount_paid" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableAmountPaid(f *float64) *MembershipUpdateOne {
	if f != nil {
		muo.SetAmountPaid(*f)
	}
	return muo
}

// AddAmountPaid adds f to the "amount_paid" field.
func (muo *MembershipUpdateOne) AddAmountPaid(f float64) *MembershipUpdateOne {
	muo.mutation.AddAmountPaid(f)
	return muo
}

// ClearAmountPaid clears the value of the "amount_paid" field.
func (muo *MembershipUpdateOne) ClearAmountPaid() *MembershipUpdateOne {
	muo.mutation.ClearAmountPaid()
	return muo
}

// SetAmountRemaining sets the "amount_remaining" field.
func (muo *MembershipUpdateOne) SetAmountRemaining(f float64) *MembershipUpdateOne {
	muo.mutation.ResetAmountRemaining()
	muo.mutation.SetAmountRemaining(f)
	return muo
}

// SetNillableAmountRemaining sets the "amount_remaining" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableAmountRemaining(f *float64) *MembershipUpdateOne {
	if f != nil {
		muo.SetAmountRemaining(*f)
	}
	return muo
}

// AddAmountRemaining adds f to the "amount_remaining" field.
func (muo *MembershipUpdateOne) AddAmountRemaining(f float64) *MembershipUpdateOne {
	muo.mutation.AddAmountRemaining(f)
	return muo
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (muo *MembershipUpdateOne) SetCustomerID(id uuid.UUID) *MembershipUpdateOne {
	muo.mutation.SetCustomerID(id)
	return muo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (muo *MembershipUpdateOne) SetCustomer(c *Customer) *MembershipUpdateOne {
	return muo.SetCustomerID(c.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (muo *MembershipUpdateOne) Mutation() *MembershipMutation {
	return muo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (muo *MembershipUpdateOne) ClearCustomer() *MembershipUpdateOne {
	muo.mutation.ClearCustomer()
	return muo
}

// Where appends a list predicates to the MembershipUpdate builder.
func (muo *MembershipUpdateOne) Where(ps ...predicate.Membership) *MembershipUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MembershipUpdateOne) Select(field string, fields ...string) *MembershipUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Membership entity.
func (muo *MembershipUpdateOne) Save(ctx context.Context) (*Membership, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MembershipUpdateOne) SaveX(ctx context.Context) *Membership {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MembershipUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MembershipUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MembershipUpdateOne) check() error {
	if _, ok := muo.mutation.CustomerID(); muo.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Membership.customer"`)
	}
	return nil
}

func (muo *MembershipUpdateOne) sqlSave(ctx context.Context) (_node *Membership, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Membership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membership.FieldID)
		for _, f := range fields {
			if !membership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != membership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.OperationDate(); ok {
		_spec.SetField(membership.FieldOperationDate, field.TypeTime, value)
	}
	if value, ok := muo.mutation.StartDate(); ok {
		_spec.SetField(membership.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := muo.mutation.EndDate(); ok {
		_spec.SetField(membership.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := muo.mutation.AmountPaid(); ok {
		_spec.SetField(membership.FieldAmountPaid, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.AddedAmountPaid(); ok {
		_spec.AddField(membership.FieldAmountPaid, field.TypeFloat64, value)
	}
	if muo.mutation.AmountPaidCleared() {
		_spec.ClearField(membership.FieldAmountPaid, field.TypeFloat64)
	}
	if value, ok := muo.mutation.AmountRemaining(); ok {
		_spec.SetField(membership.FieldAmountRemaining, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.AddedAmountRemaining(); ok {
		_spec.AddField(membership.FieldAmountRemaining, field.TypeFloat64, value)
	}
	if muo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.CustomerTable,
			Columns: []string{membership.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.CustomerTable,
			Columns: []string{membership.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Membership{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}