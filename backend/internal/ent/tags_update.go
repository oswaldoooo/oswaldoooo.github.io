// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/internal/ent/predicate"
	"backend/internal/ent/tags"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagsUpdate is the builder for updating Tags entities.
type TagsUpdate struct {
	config
	hooks     []Hook
	mutation  *TagsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TagsUpdate builder.
func (tu *TagsUpdate) Where(ps ...predicate.Tags) *TagsUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TagsUpdate) SetName(s string) *TagsUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableName(s *string) *TagsUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TagsUpdate) SetUserID(u uint64) *TagsUpdate {
	tu.mutation.ResetUserID()
	tu.mutation.SetUserID(u)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableUserID(u *uint64) *TagsUpdate {
	if u != nil {
		tu.SetUserID(*u)
	}
	return tu
}

// AddUserID adds u to the "user_id" field.
func (tu *TagsUpdate) AddUserID(u int64) *TagsUpdate {
	tu.mutation.AddUserID(u)
	return tu
}

// SetArticleID sets the "article_id" field.
func (tu *TagsUpdate) SetArticleID(u uint64) *TagsUpdate {
	tu.mutation.ResetArticleID()
	tu.mutation.SetArticleID(u)
	return tu
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableArticleID(u *uint64) *TagsUpdate {
	if u != nil {
		tu.SetArticleID(*u)
	}
	return tu
}

// AddArticleID adds u to the "article_id" field.
func (tu *TagsUpdate) AddArticleID(u int64) *TagsUpdate {
	tu.mutation.AddArticleID(u)
	return tu
}

// Mutation returns the TagsMutation object of the builder.
func (tu *TagsUpdate) Mutation() *TagsMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagsUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagsUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagsUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TagsUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TagsUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TagsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tags.Table, tags.Columns, sqlgraph.NewFieldSpec(tags.FieldID, field.TypeUint64))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tags.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.UserID(); ok {
		_spec.SetField(tags.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.AddedUserID(); ok {
		_spec.AddField(tags.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.ArticleID(); ok {
		_spec.SetField(tags.FieldArticleID, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.AddedArticleID(); ok {
		_spec.AddField(tags.FieldArticleID, field.TypeUint64, value)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tags.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagsUpdateOne is the builder for updating a single Tags entity.
type TagsUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TagsMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (tuo *TagsUpdateOne) SetName(s string) *TagsUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableName(s *string) *TagsUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TagsUpdateOne) SetUserID(u uint64) *TagsUpdateOne {
	tuo.mutation.ResetUserID()
	tuo.mutation.SetUserID(u)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableUserID(u *uint64) *TagsUpdateOne {
	if u != nil {
		tuo.SetUserID(*u)
	}
	return tuo
}

// AddUserID adds u to the "user_id" field.
func (tuo *TagsUpdateOne) AddUserID(u int64) *TagsUpdateOne {
	tuo.mutation.AddUserID(u)
	return tuo
}

// SetArticleID sets the "article_id" field.
func (tuo *TagsUpdateOne) SetArticleID(u uint64) *TagsUpdateOne {
	tuo.mutation.ResetArticleID()
	tuo.mutation.SetArticleID(u)
	return tuo
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableArticleID(u *uint64) *TagsUpdateOne {
	if u != nil {
		tuo.SetArticleID(*u)
	}
	return tuo
}

// AddArticleID adds u to the "article_id" field.
func (tuo *TagsUpdateOne) AddArticleID(u int64) *TagsUpdateOne {
	tuo.mutation.AddArticleID(u)
	return tuo
}

// Mutation returns the TagsMutation object of the builder.
func (tuo *TagsUpdateOne) Mutation() *TagsMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TagsUpdate builder.
func (tuo *TagsUpdateOne) Where(ps ...predicate.Tags) *TagsUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagsUpdateOne) Select(field string, fields ...string) *TagsUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tags entity.
func (tuo *TagsUpdateOne) Save(ctx context.Context) (*Tags, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagsUpdateOne) SaveX(ctx context.Context) *Tags {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagsUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagsUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TagsUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TagsUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TagsUpdateOne) sqlSave(ctx context.Context) (_node *Tags, err error) {
	_spec := sqlgraph.NewUpdateSpec(tags.Table, tags.Columns, sqlgraph.NewFieldSpec(tags.FieldID, field.TypeUint64))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tags.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tags.FieldID)
		for _, f := range fields {
			if !tags.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tags.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tags.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.UserID(); ok {
		_spec.SetField(tags.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.AddedUserID(); ok {
		_spec.AddField(tags.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.ArticleID(); ok {
		_spec.SetField(tags.FieldArticleID, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.AddedArticleID(); ok {
		_spec.AddField(tags.FieldArticleID, field.TypeUint64, value)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Tags{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tags.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}