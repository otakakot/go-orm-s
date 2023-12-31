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
	"github.com/otakakot/go-orm-s/gen/ent/entuser"
	"github.com/otakakot/go-orm-s/gen/ent/entusername"
	"github.com/otakakot/go-orm-s/gen/ent/predicate"
)

// EntUserNameUpdate is the builder for updating EntUserName entities.
type EntUserNameUpdate struct {
	config
	hooks    []Hook
	mutation *EntUserNameMutation
}

// Where appends a list predicates to the EntUserNameUpdate builder.
func (eunu *EntUserNameUpdate) Where(ps ...predicate.EntUserName) *EntUserNameUpdate {
	eunu.mutation.Where(ps...)
	return eunu
}

// SetUserID sets the "user_id" field.
func (eunu *EntUserNameUpdate) SetUserID(u uuid.UUID) *EntUserNameUpdate {
	eunu.mutation.SetUserID(u)
	return eunu
}

// SetValue sets the "value" field.
func (eunu *EntUserNameUpdate) SetValue(s string) *EntUserNameUpdate {
	eunu.mutation.SetValue(s)
	return eunu
}

// SetCreatedAt sets the "created_at" field.
func (eunu *EntUserNameUpdate) SetCreatedAt(t time.Time) *EntUserNameUpdate {
	eunu.mutation.SetCreatedAt(t)
	return eunu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eunu *EntUserNameUpdate) SetNillableCreatedAt(t *time.Time) *EntUserNameUpdate {
	if t != nil {
		eunu.SetCreatedAt(*t)
	}
	return eunu
}

// SetUpdatedAt sets the "updated_at" field.
func (eunu *EntUserNameUpdate) SetUpdatedAt(t time.Time) *EntUserNameUpdate {
	eunu.mutation.SetUpdatedAt(t)
	return eunu
}

// SetDeleted sets the "deleted" field.
func (eunu *EntUserNameUpdate) SetDeleted(b bool) *EntUserNameUpdate {
	eunu.mutation.SetDeleted(b)
	return eunu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (eunu *EntUserNameUpdate) SetNillableDeleted(b *bool) *EntUserNameUpdate {
	if b != nil {
		eunu.SetDeleted(*b)
	}
	return eunu
}

// SetEntUserID sets the "ent_user" edge to the EntUser entity by ID.
func (eunu *EntUserNameUpdate) SetEntUserID(id uuid.UUID) *EntUserNameUpdate {
	eunu.mutation.SetEntUserID(id)
	return eunu
}

// SetEntUser sets the "ent_user" edge to the EntUser entity.
func (eunu *EntUserNameUpdate) SetEntUser(e *EntUser) *EntUserNameUpdate {
	return eunu.SetEntUserID(e.ID)
}

// Mutation returns the EntUserNameMutation object of the builder.
func (eunu *EntUserNameUpdate) Mutation() *EntUserNameMutation {
	return eunu.mutation
}

// ClearEntUser clears the "ent_user" edge to the EntUser entity.
func (eunu *EntUserNameUpdate) ClearEntUser() *EntUserNameUpdate {
	eunu.mutation.ClearEntUser()
	return eunu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eunu *EntUserNameUpdate) Save(ctx context.Context) (int, error) {
	eunu.defaults()
	return withHooks(ctx, eunu.sqlSave, eunu.mutation, eunu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eunu *EntUserNameUpdate) SaveX(ctx context.Context) int {
	affected, err := eunu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eunu *EntUserNameUpdate) Exec(ctx context.Context) error {
	_, err := eunu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eunu *EntUserNameUpdate) ExecX(ctx context.Context) {
	if err := eunu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eunu *EntUserNameUpdate) defaults() {
	if _, ok := eunu.mutation.UpdatedAt(); !ok {
		v := entusername.UpdateDefaultUpdatedAt()
		eunu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eunu *EntUserNameUpdate) check() error {
	if _, ok := eunu.mutation.EntUserID(); eunu.mutation.EntUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EntUserName.ent_user"`)
	}
	return nil
}

func (eunu *EntUserNameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eunu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(entusername.Table, entusername.Columns, sqlgraph.NewFieldSpec(entusername.FieldID, field.TypeUUID))
	if ps := eunu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eunu.mutation.Value(); ok {
		_spec.SetField(entusername.FieldValue, field.TypeString, value)
	}
	if value, ok := eunu.mutation.CreatedAt(); ok {
		_spec.SetField(entusername.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := eunu.mutation.UpdatedAt(); ok {
		_spec.SetField(entusername.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eunu.mutation.Deleted(); ok {
		_spec.SetField(entusername.FieldDeleted, field.TypeBool, value)
	}
	if eunu.mutation.EntUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entusername.EntUserTable,
			Columns: []string{entusername.EntUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entuser.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eunu.mutation.EntUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entusername.EntUserTable,
			Columns: []string{entusername.EntUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entuser.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eunu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entusername.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eunu.mutation.done = true
	return n, nil
}

// EntUserNameUpdateOne is the builder for updating a single EntUserName entity.
type EntUserNameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntUserNameMutation
}

// SetUserID sets the "user_id" field.
func (eunuo *EntUserNameUpdateOne) SetUserID(u uuid.UUID) *EntUserNameUpdateOne {
	eunuo.mutation.SetUserID(u)
	return eunuo
}

// SetValue sets the "value" field.
func (eunuo *EntUserNameUpdateOne) SetValue(s string) *EntUserNameUpdateOne {
	eunuo.mutation.SetValue(s)
	return eunuo
}

// SetCreatedAt sets the "created_at" field.
func (eunuo *EntUserNameUpdateOne) SetCreatedAt(t time.Time) *EntUserNameUpdateOne {
	eunuo.mutation.SetCreatedAt(t)
	return eunuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eunuo *EntUserNameUpdateOne) SetNillableCreatedAt(t *time.Time) *EntUserNameUpdateOne {
	if t != nil {
		eunuo.SetCreatedAt(*t)
	}
	return eunuo
}

// SetUpdatedAt sets the "updated_at" field.
func (eunuo *EntUserNameUpdateOne) SetUpdatedAt(t time.Time) *EntUserNameUpdateOne {
	eunuo.mutation.SetUpdatedAt(t)
	return eunuo
}

// SetDeleted sets the "deleted" field.
func (eunuo *EntUserNameUpdateOne) SetDeleted(b bool) *EntUserNameUpdateOne {
	eunuo.mutation.SetDeleted(b)
	return eunuo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (eunuo *EntUserNameUpdateOne) SetNillableDeleted(b *bool) *EntUserNameUpdateOne {
	if b != nil {
		eunuo.SetDeleted(*b)
	}
	return eunuo
}

// SetEntUserID sets the "ent_user" edge to the EntUser entity by ID.
func (eunuo *EntUserNameUpdateOne) SetEntUserID(id uuid.UUID) *EntUserNameUpdateOne {
	eunuo.mutation.SetEntUserID(id)
	return eunuo
}

// SetEntUser sets the "ent_user" edge to the EntUser entity.
func (eunuo *EntUserNameUpdateOne) SetEntUser(e *EntUser) *EntUserNameUpdateOne {
	return eunuo.SetEntUserID(e.ID)
}

// Mutation returns the EntUserNameMutation object of the builder.
func (eunuo *EntUserNameUpdateOne) Mutation() *EntUserNameMutation {
	return eunuo.mutation
}

// ClearEntUser clears the "ent_user" edge to the EntUser entity.
func (eunuo *EntUserNameUpdateOne) ClearEntUser() *EntUserNameUpdateOne {
	eunuo.mutation.ClearEntUser()
	return eunuo
}

// Where appends a list predicates to the EntUserNameUpdate builder.
func (eunuo *EntUserNameUpdateOne) Where(ps ...predicate.EntUserName) *EntUserNameUpdateOne {
	eunuo.mutation.Where(ps...)
	return eunuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eunuo *EntUserNameUpdateOne) Select(field string, fields ...string) *EntUserNameUpdateOne {
	eunuo.fields = append([]string{field}, fields...)
	return eunuo
}

// Save executes the query and returns the updated EntUserName entity.
func (eunuo *EntUserNameUpdateOne) Save(ctx context.Context) (*EntUserName, error) {
	eunuo.defaults()
	return withHooks(ctx, eunuo.sqlSave, eunuo.mutation, eunuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eunuo *EntUserNameUpdateOne) SaveX(ctx context.Context) *EntUserName {
	node, err := eunuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eunuo *EntUserNameUpdateOne) Exec(ctx context.Context) error {
	_, err := eunuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eunuo *EntUserNameUpdateOne) ExecX(ctx context.Context) {
	if err := eunuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eunuo *EntUserNameUpdateOne) defaults() {
	if _, ok := eunuo.mutation.UpdatedAt(); !ok {
		v := entusername.UpdateDefaultUpdatedAt()
		eunuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eunuo *EntUserNameUpdateOne) check() error {
	if _, ok := eunuo.mutation.EntUserID(); eunuo.mutation.EntUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EntUserName.ent_user"`)
	}
	return nil
}

func (eunuo *EntUserNameUpdateOne) sqlSave(ctx context.Context) (_node *EntUserName, err error) {
	if err := eunuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(entusername.Table, entusername.Columns, sqlgraph.NewFieldSpec(entusername.FieldID, field.TypeUUID))
	id, ok := eunuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntUserName.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eunuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entusername.FieldID)
		for _, f := range fields {
			if !entusername.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entusername.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eunuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eunuo.mutation.Value(); ok {
		_spec.SetField(entusername.FieldValue, field.TypeString, value)
	}
	if value, ok := eunuo.mutation.CreatedAt(); ok {
		_spec.SetField(entusername.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := eunuo.mutation.UpdatedAt(); ok {
		_spec.SetField(entusername.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eunuo.mutation.Deleted(); ok {
		_spec.SetField(entusername.FieldDeleted, field.TypeBool, value)
	}
	if eunuo.mutation.EntUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entusername.EntUserTable,
			Columns: []string{entusername.EntUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entuser.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eunuo.mutation.EntUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entusername.EntUserTable,
			Columns: []string{entusername.EntUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entuser.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntUserName{config: eunuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eunuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entusername.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eunuo.mutation.done = true
	return _node, nil
}
