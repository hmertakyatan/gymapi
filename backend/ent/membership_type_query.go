// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent/membership_type"
	"github.com/hmertakyatan/gymapi/ent/predicate"
)

// MembershipTypeQuery is the builder for querying Membership_type entities.
type MembershipTypeQuery struct {
	config
	ctx        *QueryContext
	order      []membership_type.OrderOption
	inters     []Interceptor
	predicates []predicate.Membership_type
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MembershipTypeQuery builder.
func (mtq *MembershipTypeQuery) Where(ps ...predicate.Membership_type) *MembershipTypeQuery {
	mtq.predicates = append(mtq.predicates, ps...)
	return mtq
}

// Limit the number of records to be returned by this query.
func (mtq *MembershipTypeQuery) Limit(limit int) *MembershipTypeQuery {
	mtq.ctx.Limit = &limit
	return mtq
}

// Offset to start from.
func (mtq *MembershipTypeQuery) Offset(offset int) *MembershipTypeQuery {
	mtq.ctx.Offset = &offset
	return mtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mtq *MembershipTypeQuery) Unique(unique bool) *MembershipTypeQuery {
	mtq.ctx.Unique = &unique
	return mtq
}

// Order specifies how the records should be ordered.
func (mtq *MembershipTypeQuery) Order(o ...membership_type.OrderOption) *MembershipTypeQuery {
	mtq.order = append(mtq.order, o...)
	return mtq
}

// First returns the first Membership_type entity from the query.
// Returns a *NotFoundError when no Membership_type was found.
func (mtq *MembershipTypeQuery) First(ctx context.Context) (*Membership_type, error) {
	nodes, err := mtq.Limit(1).All(setContextOp(ctx, mtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{membership_type.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mtq *MembershipTypeQuery) FirstX(ctx context.Context) *Membership_type {
	node, err := mtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Membership_type ID from the query.
// Returns a *NotFoundError when no Membership_type ID was found.
func (mtq *MembershipTypeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mtq.Limit(1).IDs(setContextOp(ctx, mtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{membership_type.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mtq *MembershipTypeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Membership_type entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Membership_type entity is found.
// Returns a *NotFoundError when no Membership_type entities are found.
func (mtq *MembershipTypeQuery) Only(ctx context.Context) (*Membership_type, error) {
	nodes, err := mtq.Limit(2).All(setContextOp(ctx, mtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{membership_type.Label}
	default:
		return nil, &NotSingularError{membership_type.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mtq *MembershipTypeQuery) OnlyX(ctx context.Context) *Membership_type {
	node, err := mtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Membership_type ID in the query.
// Returns a *NotSingularError when more than one Membership_type ID is found.
// Returns a *NotFoundError when no entities are found.
func (mtq *MembershipTypeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mtq.Limit(2).IDs(setContextOp(ctx, mtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{membership_type.Label}
	default:
		err = &NotSingularError{membership_type.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mtq *MembershipTypeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Membership_types.
func (mtq *MembershipTypeQuery) All(ctx context.Context) ([]*Membership_type, error) {
	ctx = setContextOp(ctx, mtq.ctx, "All")
	if err := mtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Membership_type, *MembershipTypeQuery]()
	return withInterceptors[[]*Membership_type](ctx, mtq, qr, mtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mtq *MembershipTypeQuery) AllX(ctx context.Context) []*Membership_type {
	nodes, err := mtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Membership_type IDs.
func (mtq *MembershipTypeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mtq.ctx.Unique == nil && mtq.path != nil {
		mtq.Unique(true)
	}
	ctx = setContextOp(ctx, mtq.ctx, "IDs")
	if err = mtq.Select(membership_type.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mtq *MembershipTypeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mtq *MembershipTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mtq.ctx, "Count")
	if err := mtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mtq, querierCount[*MembershipTypeQuery](), mtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mtq *MembershipTypeQuery) CountX(ctx context.Context) int {
	count, err := mtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mtq *MembershipTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mtq.ctx, "Exist")
	switch _, err := mtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mtq *MembershipTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := mtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MembershipTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mtq *MembershipTypeQuery) Clone() *MembershipTypeQuery {
	if mtq == nil {
		return nil
	}
	return &MembershipTypeQuery{
		config:     mtq.config,
		ctx:        mtq.ctx.Clone(),
		order:      append([]membership_type.OrderOption{}, mtq.order...),
		inters:     append([]Interceptor{}, mtq.inters...),
		predicates: append([]predicate.Membership_type{}, mtq.predicates...),
		// clone intermediate query.
		sql:  mtq.sql.Clone(),
		path: mtq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MembershipType.Query().
//		GroupBy(membership_type.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mtq *MembershipTypeQuery) GroupBy(field string, fields ...string) *MembershipTypeGroupBy {
	mtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MembershipTypeGroupBy{build: mtq}
	grbuild.flds = &mtq.ctx.Fields
	grbuild.label = membership_type.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MembershipType.Query().
//		Select(membership_type.FieldName).
//		Scan(ctx, &v)
func (mtq *MembershipTypeQuery) Select(fields ...string) *MembershipTypeSelect {
	mtq.ctx.Fields = append(mtq.ctx.Fields, fields...)
	sbuild := &MembershipTypeSelect{MembershipTypeQuery: mtq}
	sbuild.label = membership_type.Label
	sbuild.flds, sbuild.scan = &mtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MembershipTypeSelect configured with the given aggregations.
func (mtq *MembershipTypeQuery) Aggregate(fns ...AggregateFunc) *MembershipTypeSelect {
	return mtq.Select().Aggregate(fns...)
}

func (mtq *MembershipTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mtq); err != nil {
				return err
			}
		}
	}
	for _, f := range mtq.ctx.Fields {
		if !membership_type.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mtq.path != nil {
		prev, err := mtq.path(ctx)
		if err != nil {
			return err
		}
		mtq.sql = prev
	}
	return nil
}

func (mtq *MembershipTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Membership_type, error) {
	var (
		nodes = []*Membership_type{}
		_spec = mtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Membership_type).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Membership_type{config: mtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mtq *MembershipTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mtq.querySpec()
	_spec.Node.Columns = mtq.ctx.Fields
	if len(mtq.ctx.Fields) > 0 {
		_spec.Unique = mtq.ctx.Unique != nil && *mtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mtq.driver, _spec)
}

func (mtq *MembershipTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(membership_type.Table, membership_type.Columns, sqlgraph.NewFieldSpec(membership_type.FieldID, field.TypeUUID))
	_spec.From = mtq.sql
	if unique := mtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mtq.path != nil {
		_spec.Unique = true
	}
	if fields := mtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membership_type.FieldID)
		for i := range fields {
			if fields[i] != membership_type.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mtq *MembershipTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mtq.driver.Dialect())
	t1 := builder.Table(membership_type.Table)
	columns := mtq.ctx.Fields
	if len(columns) == 0 {
		columns = membership_type.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mtq.sql != nil {
		selector = mtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mtq.ctx.Unique != nil && *mtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mtq.predicates {
		p(selector)
	}
	for _, p := range mtq.order {
		p(selector)
	}
	if offset := mtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MembershipTypeGroupBy is the group-by builder for Membership_type entities.
type MembershipTypeGroupBy struct {
	selector
	build *MembershipTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mtgb *MembershipTypeGroupBy) Aggregate(fns ...AggregateFunc) *MembershipTypeGroupBy {
	mtgb.fns = append(mtgb.fns, fns...)
	return mtgb
}

// Scan applies the selector query and scans the result into the given value.
func (mtgb *MembershipTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mtgb.build.ctx, "GroupBy")
	if err := mtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MembershipTypeQuery, *MembershipTypeGroupBy](ctx, mtgb.build, mtgb, mtgb.build.inters, v)
}

func (mtgb *MembershipTypeGroupBy) sqlScan(ctx context.Context, root *MembershipTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mtgb.fns))
	for _, fn := range mtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mtgb.flds)+len(mtgb.fns))
		for _, f := range *mtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MembershipTypeSelect is the builder for selecting fields of MembershipType entities.
type MembershipTypeSelect struct {
	*MembershipTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mts *MembershipTypeSelect) Aggregate(fns ...AggregateFunc) *MembershipTypeSelect {
	mts.fns = append(mts.fns, fns...)
	return mts
}

// Scan applies the selector query and scans the result into the given value.
func (mts *MembershipTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mts.ctx, "Select")
	if err := mts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MembershipTypeQuery, *MembershipTypeSelect](ctx, mts.MembershipTypeQuery, mts, mts.inters, v)
}

func (mts *MembershipTypeSelect) sqlScan(ctx context.Context, root *MembershipTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mts.fns))
	for _, fn := range mts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}