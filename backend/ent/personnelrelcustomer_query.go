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
	"github.com/hmertakyatan/gymapi/ent/customer"
	"github.com/hmertakyatan/gymapi/ent/personnel"
	"github.com/hmertakyatan/gymapi/ent/personnelrelcustomer"
	"github.com/hmertakyatan/gymapi/ent/predicate"
)

// PersonnelRelCustomerQuery is the builder for querying PersonnelRelCustomer entities.
type PersonnelRelCustomerQuery struct {
	config
	ctx           *QueryContext
	order         []personnelrelcustomer.OrderOption
	inters        []Interceptor
	predicates    []predicate.PersonnelRelCustomer
	withPersonnel *PersonnelQuery
	withCustomer  *CustomerQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PersonnelRelCustomerQuery builder.
func (prcq *PersonnelRelCustomerQuery) Where(ps ...predicate.PersonnelRelCustomer) *PersonnelRelCustomerQuery {
	prcq.predicates = append(prcq.predicates, ps...)
	return prcq
}

// Limit the number of records to be returned by this query.
func (prcq *PersonnelRelCustomerQuery) Limit(limit int) *PersonnelRelCustomerQuery {
	prcq.ctx.Limit = &limit
	return prcq
}

// Offset to start from.
func (prcq *PersonnelRelCustomerQuery) Offset(offset int) *PersonnelRelCustomerQuery {
	prcq.ctx.Offset = &offset
	return prcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prcq *PersonnelRelCustomerQuery) Unique(unique bool) *PersonnelRelCustomerQuery {
	prcq.ctx.Unique = &unique
	return prcq
}

// Order specifies how the records should be ordered.
func (prcq *PersonnelRelCustomerQuery) Order(o ...personnelrelcustomer.OrderOption) *PersonnelRelCustomerQuery {
	prcq.order = append(prcq.order, o...)
	return prcq
}

// QueryPersonnel chains the current query on the "personnel" edge.
func (prcq *PersonnelRelCustomerQuery) QueryPersonnel() *PersonnelQuery {
	query := (&PersonnelClient{config: prcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(personnelrelcustomer.Table, personnelrelcustomer.FieldID, selector),
			sqlgraph.To(personnel.Table, personnel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, personnelrelcustomer.PersonnelTable, personnelrelcustomer.PersonnelColumn),
		)
		fromU = sqlgraph.SetNeighbors(prcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomer chains the current query on the "customer" edge.
func (prcq *PersonnelRelCustomerQuery) QueryCustomer() *CustomerQuery {
	query := (&CustomerClient{config: prcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(personnelrelcustomer.Table, personnelrelcustomer.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, personnelrelcustomer.CustomerTable, personnelrelcustomer.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(prcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PersonnelRelCustomer entity from the query.
// Returns a *NotFoundError when no PersonnelRelCustomer was found.
func (prcq *PersonnelRelCustomerQuery) First(ctx context.Context) (*PersonnelRelCustomer, error) {
	nodes, err := prcq.Limit(1).All(setContextOp(ctx, prcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{personnelrelcustomer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) FirstX(ctx context.Context) *PersonnelRelCustomer {
	node, err := prcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PersonnelRelCustomer ID from the query.
// Returns a *NotFoundError when no PersonnelRelCustomer ID was found.
func (prcq *PersonnelRelCustomerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = prcq.Limit(1).IDs(setContextOp(ctx, prcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{personnelrelcustomer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := prcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PersonnelRelCustomer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PersonnelRelCustomer entity is found.
// Returns a *NotFoundError when no PersonnelRelCustomer entities are found.
func (prcq *PersonnelRelCustomerQuery) Only(ctx context.Context) (*PersonnelRelCustomer, error) {
	nodes, err := prcq.Limit(2).All(setContextOp(ctx, prcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{personnelrelcustomer.Label}
	default:
		return nil, &NotSingularError{personnelrelcustomer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) OnlyX(ctx context.Context) *PersonnelRelCustomer {
	node, err := prcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PersonnelRelCustomer ID in the query.
// Returns a *NotSingularError when more than one PersonnelRelCustomer ID is found.
// Returns a *NotFoundError when no entities are found.
func (prcq *PersonnelRelCustomerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = prcq.Limit(2).IDs(setContextOp(ctx, prcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{personnelrelcustomer.Label}
	default:
		err = &NotSingularError{personnelrelcustomer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := prcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PersonnelRelCustomers.
func (prcq *PersonnelRelCustomerQuery) All(ctx context.Context) ([]*PersonnelRelCustomer, error) {
	ctx = setContextOp(ctx, prcq.ctx, "All")
	if err := prcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PersonnelRelCustomer, *PersonnelRelCustomerQuery]()
	return withInterceptors[[]*PersonnelRelCustomer](ctx, prcq, qr, prcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) AllX(ctx context.Context) []*PersonnelRelCustomer {
	nodes, err := prcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PersonnelRelCustomer IDs.
func (prcq *PersonnelRelCustomerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if prcq.ctx.Unique == nil && prcq.path != nil {
		prcq.Unique(true)
	}
	ctx = setContextOp(ctx, prcq.ctx, "IDs")
	if err = prcq.Select(personnelrelcustomer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := prcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prcq *PersonnelRelCustomerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, prcq.ctx, "Count")
	if err := prcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, prcq, querierCount[*PersonnelRelCustomerQuery](), prcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) CountX(ctx context.Context) int {
	count, err := prcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prcq *PersonnelRelCustomerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, prcq.ctx, "Exist")
	switch _, err := prcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (prcq *PersonnelRelCustomerQuery) ExistX(ctx context.Context) bool {
	exist, err := prcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PersonnelRelCustomerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prcq *PersonnelRelCustomerQuery) Clone() *PersonnelRelCustomerQuery {
	if prcq == nil {
		return nil
	}
	return &PersonnelRelCustomerQuery{
		config:        prcq.config,
		ctx:           prcq.ctx.Clone(),
		order:         append([]personnelrelcustomer.OrderOption{}, prcq.order...),
		inters:        append([]Interceptor{}, prcq.inters...),
		predicates:    append([]predicate.PersonnelRelCustomer{}, prcq.predicates...),
		withPersonnel: prcq.withPersonnel.Clone(),
		withCustomer:  prcq.withCustomer.Clone(),
		// clone intermediate query.
		sql:  prcq.sql.Clone(),
		path: prcq.path,
	}
}

// WithPersonnel tells the query-builder to eager-load the nodes that are connected to
// the "personnel" edge. The optional arguments are used to configure the query builder of the edge.
func (prcq *PersonnelRelCustomerQuery) WithPersonnel(opts ...func(*PersonnelQuery)) *PersonnelRelCustomerQuery {
	query := (&PersonnelClient{config: prcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prcq.withPersonnel = query
	return prcq
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (prcq *PersonnelRelCustomerQuery) WithCustomer(opts ...func(*CustomerQuery)) *PersonnelRelCustomerQuery {
	query := (&CustomerClient{config: prcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prcq.withCustomer = query
	return prcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PersonnelRelCustomer.Query().
//		GroupBy(personnelrelcustomer.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (prcq *PersonnelRelCustomerQuery) GroupBy(field string, fields ...string) *PersonnelRelCustomerGroupBy {
	prcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PersonnelRelCustomerGroupBy{build: prcq}
	grbuild.flds = &prcq.ctx.Fields
	grbuild.label = personnelrelcustomer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.PersonnelRelCustomer.Query().
//		Select(personnelrelcustomer.FieldDescription).
//		Scan(ctx, &v)
func (prcq *PersonnelRelCustomerQuery) Select(fields ...string) *PersonnelRelCustomerSelect {
	prcq.ctx.Fields = append(prcq.ctx.Fields, fields...)
	sbuild := &PersonnelRelCustomerSelect{PersonnelRelCustomerQuery: prcq}
	sbuild.label = personnelrelcustomer.Label
	sbuild.flds, sbuild.scan = &prcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PersonnelRelCustomerSelect configured with the given aggregations.
func (prcq *PersonnelRelCustomerQuery) Aggregate(fns ...AggregateFunc) *PersonnelRelCustomerSelect {
	return prcq.Select().Aggregate(fns...)
}

func (prcq *PersonnelRelCustomerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range prcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, prcq); err != nil {
				return err
			}
		}
	}
	for _, f := range prcq.ctx.Fields {
		if !personnelrelcustomer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prcq.path != nil {
		prev, err := prcq.path(ctx)
		if err != nil {
			return err
		}
		prcq.sql = prev
	}
	return nil
}

func (prcq *PersonnelRelCustomerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PersonnelRelCustomer, error) {
	var (
		nodes       = []*PersonnelRelCustomer{}
		withFKs     = prcq.withFKs
		_spec       = prcq.querySpec()
		loadedTypes = [2]bool{
			prcq.withPersonnel != nil,
			prcq.withCustomer != nil,
		}
	)
	if prcq.withPersonnel != nil || prcq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, personnelrelcustomer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PersonnelRelCustomer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PersonnelRelCustomer{config: prcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := prcq.withPersonnel; query != nil {
		if err := prcq.loadPersonnel(ctx, query, nodes, nil,
			func(n *PersonnelRelCustomer, e *Personnel) { n.Edges.Personnel = e }); err != nil {
			return nil, err
		}
	}
	if query := prcq.withCustomer; query != nil {
		if err := prcq.loadCustomer(ctx, query, nodes, nil,
			func(n *PersonnelRelCustomer, e *Customer) { n.Edges.Customer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (prcq *PersonnelRelCustomerQuery) loadPersonnel(ctx context.Context, query *PersonnelQuery, nodes []*PersonnelRelCustomer, init func(*PersonnelRelCustomer), assign func(*PersonnelRelCustomer, *Personnel)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PersonnelRelCustomer)
	for i := range nodes {
		if nodes[i].personnel_personnel_rel_customer == nil {
			continue
		}
		fk := *nodes[i].personnel_personnel_rel_customer
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(personnel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "personnel_personnel_rel_customer" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (prcq *PersonnelRelCustomerQuery) loadCustomer(ctx context.Context, query *CustomerQuery, nodes []*PersonnelRelCustomer, init func(*PersonnelRelCustomer), assign func(*PersonnelRelCustomer, *Customer)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PersonnelRelCustomer)
	for i := range nodes {
		if nodes[i].customer_personnel_rel_customer == nil {
			continue
		}
		fk := *nodes[i].customer_personnel_rel_customer
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(customer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_personnel_rel_customer" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (prcq *PersonnelRelCustomerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prcq.querySpec()
	_spec.Node.Columns = prcq.ctx.Fields
	if len(prcq.ctx.Fields) > 0 {
		_spec.Unique = prcq.ctx.Unique != nil && *prcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, prcq.driver, _spec)
}

func (prcq *PersonnelRelCustomerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(personnelrelcustomer.Table, personnelrelcustomer.Columns, sqlgraph.NewFieldSpec(personnelrelcustomer.FieldID, field.TypeUUID))
	_spec.From = prcq.sql
	if unique := prcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if prcq.path != nil {
		_spec.Unique = true
	}
	if fields := prcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, personnelrelcustomer.FieldID)
		for i := range fields {
			if fields[i] != personnelrelcustomer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prcq *PersonnelRelCustomerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prcq.driver.Dialect())
	t1 := builder.Table(personnelrelcustomer.Table)
	columns := prcq.ctx.Fields
	if len(columns) == 0 {
		columns = personnelrelcustomer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prcq.sql != nil {
		selector = prcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prcq.ctx.Unique != nil && *prcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range prcq.predicates {
		p(selector)
	}
	for _, p := range prcq.order {
		p(selector)
	}
	if offset := prcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PersonnelRelCustomerGroupBy is the group-by builder for PersonnelRelCustomer entities.
type PersonnelRelCustomerGroupBy struct {
	selector
	build *PersonnelRelCustomerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prcgb *PersonnelRelCustomerGroupBy) Aggregate(fns ...AggregateFunc) *PersonnelRelCustomerGroupBy {
	prcgb.fns = append(prcgb.fns, fns...)
	return prcgb
}

// Scan applies the selector query and scans the result into the given value.
func (prcgb *PersonnelRelCustomerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prcgb.build.ctx, "GroupBy")
	if err := prcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonnelRelCustomerQuery, *PersonnelRelCustomerGroupBy](ctx, prcgb.build, prcgb, prcgb.build.inters, v)
}

func (prcgb *PersonnelRelCustomerGroupBy) sqlScan(ctx context.Context, root *PersonnelRelCustomerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(prcgb.fns))
	for _, fn := range prcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*prcgb.flds)+len(prcgb.fns))
		for _, f := range *prcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*prcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PersonnelRelCustomerSelect is the builder for selecting fields of PersonnelRelCustomer entities.
type PersonnelRelCustomerSelect struct {
	*PersonnelRelCustomerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (prcs *PersonnelRelCustomerSelect) Aggregate(fns ...AggregateFunc) *PersonnelRelCustomerSelect {
	prcs.fns = append(prcs.fns, fns...)
	return prcs
}

// Scan applies the selector query and scans the result into the given value.
func (prcs *PersonnelRelCustomerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prcs.ctx, "Select")
	if err := prcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonnelRelCustomerQuery, *PersonnelRelCustomerSelect](ctx, prcs.PersonnelRelCustomerQuery, prcs, prcs.inters, v)
}

func (prcs *PersonnelRelCustomerSelect) sqlScan(ctx context.Context, root *PersonnelRelCustomerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(prcs.fns))
	for _, fn := range prcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*prcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}