// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/otakakot/go-orm-s/gen/ent/entuser"
	"github.com/otakakot/go-orm-s/gen/ent/entusername"
	"github.com/otakakot/go-orm-s/gen/ent/predicate"
)

// EntUserQuery is the builder for querying EntUser entities.
type EntUserQuery struct {
	config
	ctx              *QueryContext
	order            []entuser.OrderOption
	inters           []Interceptor
	predicates       []predicate.EntUser
	withEntUserNames *EntUserNameQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntUserQuery builder.
func (euq *EntUserQuery) Where(ps ...predicate.EntUser) *EntUserQuery {
	euq.predicates = append(euq.predicates, ps...)
	return euq
}

// Limit the number of records to be returned by this query.
func (euq *EntUserQuery) Limit(limit int) *EntUserQuery {
	euq.ctx.Limit = &limit
	return euq
}

// Offset to start from.
func (euq *EntUserQuery) Offset(offset int) *EntUserQuery {
	euq.ctx.Offset = &offset
	return euq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (euq *EntUserQuery) Unique(unique bool) *EntUserQuery {
	euq.ctx.Unique = &unique
	return euq
}

// Order specifies how the records should be ordered.
func (euq *EntUserQuery) Order(o ...entuser.OrderOption) *EntUserQuery {
	euq.order = append(euq.order, o...)
	return euq
}

// QueryEntUserNames chains the current query on the "ent_user_names" edge.
func (euq *EntUserQuery) QueryEntUserNames() *EntUserNameQuery {
	query := (&EntUserNameClient{config: euq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := euq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := euq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entuser.Table, entuser.FieldID, selector),
			sqlgraph.To(entusername.Table, entusername.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, entuser.EntUserNamesTable, entuser.EntUserNamesColumn),
		)
		fromU = sqlgraph.SetNeighbors(euq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntUser entity from the query.
// Returns a *NotFoundError when no EntUser was found.
func (euq *EntUserQuery) First(ctx context.Context) (*EntUser, error) {
	nodes, err := euq.Limit(1).All(setContextOp(ctx, euq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (euq *EntUserQuery) FirstX(ctx context.Context) *EntUser {
	node, err := euq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntUser ID from the query.
// Returns a *NotFoundError when no EntUser ID was found.
func (euq *EntUserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = euq.Limit(1).IDs(setContextOp(ctx, euq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (euq *EntUserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := euq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntUser entity is found.
// Returns a *NotFoundError when no EntUser entities are found.
func (euq *EntUserQuery) Only(ctx context.Context) (*EntUser, error) {
	nodes, err := euq.Limit(2).All(setContextOp(ctx, euq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entuser.Label}
	default:
		return nil, &NotSingularError{entuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (euq *EntUserQuery) OnlyX(ctx context.Context) *EntUser {
	node, err := euq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntUser ID in the query.
// Returns a *NotSingularError when more than one EntUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (euq *EntUserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = euq.Limit(2).IDs(setContextOp(ctx, euq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entuser.Label}
	default:
		err = &NotSingularError{entuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (euq *EntUserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := euq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntUsers.
func (euq *EntUserQuery) All(ctx context.Context) ([]*EntUser, error) {
	ctx = setContextOp(ctx, euq.ctx, "All")
	if err := euq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EntUser, *EntUserQuery]()
	return withInterceptors[[]*EntUser](ctx, euq, qr, euq.inters)
}

// AllX is like All, but panics if an error occurs.
func (euq *EntUserQuery) AllX(ctx context.Context) []*EntUser {
	nodes, err := euq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntUser IDs.
func (euq *EntUserQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if euq.ctx.Unique == nil && euq.path != nil {
		euq.Unique(true)
	}
	ctx = setContextOp(ctx, euq.ctx, "IDs")
	if err = euq.Select(entuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (euq *EntUserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := euq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (euq *EntUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, euq.ctx, "Count")
	if err := euq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, euq, querierCount[*EntUserQuery](), euq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (euq *EntUserQuery) CountX(ctx context.Context) int {
	count, err := euq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (euq *EntUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, euq.ctx, "Exist")
	switch _, err := euq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (euq *EntUserQuery) ExistX(ctx context.Context) bool {
	exist, err := euq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (euq *EntUserQuery) Clone() *EntUserQuery {
	if euq == nil {
		return nil
	}
	return &EntUserQuery{
		config:           euq.config,
		ctx:              euq.ctx.Clone(),
		order:            append([]entuser.OrderOption{}, euq.order...),
		inters:           append([]Interceptor{}, euq.inters...),
		predicates:       append([]predicate.EntUser{}, euq.predicates...),
		withEntUserNames: euq.withEntUserNames.Clone(),
		// clone intermediate query.
		sql:  euq.sql.Clone(),
		path: euq.path,
	}
}

// WithEntUserNames tells the query-builder to eager-load the nodes that are connected to
// the "ent_user_names" edge. The optional arguments are used to configure the query builder of the edge.
func (euq *EntUserQuery) WithEntUserNames(opts ...func(*EntUserNameQuery)) *EntUserQuery {
	query := (&EntUserNameClient{config: euq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	euq.withEntUserNames = query
	return euq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EntUser.Query().
//		GroupBy(entuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (euq *EntUserQuery) GroupBy(field string, fields ...string) *EntUserGroupBy {
	euq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EntUserGroupBy{build: euq}
	grbuild.flds = &euq.ctx.Fields
	grbuild.label = entuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.EntUser.Query().
//		Select(entuser.FieldCreatedAt).
//		Scan(ctx, &v)
func (euq *EntUserQuery) Select(fields ...string) *EntUserSelect {
	euq.ctx.Fields = append(euq.ctx.Fields, fields...)
	sbuild := &EntUserSelect{EntUserQuery: euq}
	sbuild.label = entuser.Label
	sbuild.flds, sbuild.scan = &euq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EntUserSelect configured with the given aggregations.
func (euq *EntUserQuery) Aggregate(fns ...AggregateFunc) *EntUserSelect {
	return euq.Select().Aggregate(fns...)
}

func (euq *EntUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range euq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, euq); err != nil {
				return err
			}
		}
	}
	for _, f := range euq.ctx.Fields {
		if !entuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if euq.path != nil {
		prev, err := euq.path(ctx)
		if err != nil {
			return err
		}
		euq.sql = prev
	}
	return nil
}

func (euq *EntUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntUser, error) {
	var (
		nodes       = []*EntUser{}
		_spec       = euq.querySpec()
		loadedTypes = [1]bool{
			euq.withEntUserNames != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntUser{config: euq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, euq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := euq.withEntUserNames; query != nil {
		if err := euq.loadEntUserNames(ctx, query, nodes,
			func(n *EntUser) { n.Edges.EntUserNames = []*EntUserName{} },
			func(n *EntUser, e *EntUserName) { n.Edges.EntUserNames = append(n.Edges.EntUserNames, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (euq *EntUserQuery) loadEntUserNames(ctx context.Context, query *EntUserNameQuery, nodes []*EntUser, init func(*EntUser), assign func(*EntUser, *EntUserName)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*EntUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(entusername.FieldUserID)
	}
	query.Where(predicate.EntUserName(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(entuser.EntUserNamesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (euq *EntUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := euq.querySpec()
	_spec.Node.Columns = euq.ctx.Fields
	if len(euq.ctx.Fields) > 0 {
		_spec.Unique = euq.ctx.Unique != nil && *euq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, euq.driver, _spec)
}

func (euq *EntUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entuser.Table, entuser.Columns, sqlgraph.NewFieldSpec(entuser.FieldID, field.TypeUUID))
	_spec.From = euq.sql
	if unique := euq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if euq.path != nil {
		_spec.Unique = true
	}
	if fields := euq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entuser.FieldID)
		for i := range fields {
			if fields[i] != entuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := euq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := euq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := euq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := euq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (euq *EntUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(euq.driver.Dialect())
	t1 := builder.Table(entuser.Table)
	columns := euq.ctx.Fields
	if len(columns) == 0 {
		columns = entuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if euq.sql != nil {
		selector = euq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if euq.ctx.Unique != nil && *euq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range euq.predicates {
		p(selector)
	}
	for _, p := range euq.order {
		p(selector)
	}
	if offset := euq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := euq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntUserGroupBy is the group-by builder for EntUser entities.
type EntUserGroupBy struct {
	selector
	build *EntUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eugb *EntUserGroupBy) Aggregate(fns ...AggregateFunc) *EntUserGroupBy {
	eugb.fns = append(eugb.fns, fns...)
	return eugb
}

// Scan applies the selector query and scans the result into the given value.
func (eugb *EntUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eugb.build.ctx, "GroupBy")
	if err := eugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntUserQuery, *EntUserGroupBy](ctx, eugb.build, eugb, eugb.build.inters, v)
}

func (eugb *EntUserGroupBy) sqlScan(ctx context.Context, root *EntUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eugb.fns))
	for _, fn := range eugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eugb.flds)+len(eugb.fns))
		for _, f := range *eugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EntUserSelect is the builder for selecting fields of EntUser entities.
type EntUserSelect struct {
	*EntUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eus *EntUserSelect) Aggregate(fns ...AggregateFunc) *EntUserSelect {
	eus.fns = append(eus.fns, fns...)
	return eus
}

// Scan applies the selector query and scans the result into the given value.
func (eus *EntUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eus.ctx, "Select")
	if err := eus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntUserQuery, *EntUserSelect](ctx, eus.EntUserQuery, eus, eus.inters, v)
}

func (eus *EntUserSelect) sqlScan(ctx context.Context, root *EntUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eus.fns))
	for _, fn := range eus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
