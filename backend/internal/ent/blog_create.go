// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/internal/ent/blog"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogCreate is the builder for creating a Blog entity.
type BlogCreate struct {
	config
	mutation *BlogMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (bc *BlogCreate) SetTitle(s string) *BlogCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetAddress sets the "address" field.
func (bc *BlogCreate) SetAddress(s string) *BlogCreate {
	bc.mutation.SetAddress(s)
	return bc
}

// SetUserID sets the "user_id" field.
func (bc *BlogCreate) SetUserID(u uint64) *BlogCreate {
	bc.mutation.SetUserID(u)
	return bc
}

// SetCtime sets the "ctime" field.
func (bc *BlogCreate) SetCtime(i int64) *BlogCreate {
	bc.mutation.SetCtime(i)
	return bc
}

// SetDesc sets the "desc" field.
func (bc *BlogCreate) SetDesc(s string) *BlogCreate {
	bc.mutation.SetDesc(s)
	return bc
}

// SetID sets the "id" field.
func (bc *BlogCreate) SetID(u uint64) *BlogCreate {
	bc.mutation.SetID(u)
	return bc
}

// Mutation returns the BlogMutation object of the builder.
func (bc *BlogCreate) Mutation() *BlogMutation {
	return bc.mutation
}

// Save creates the Blog in the database.
func (bc *BlogCreate) Save(ctx context.Context) (*Blog, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlogCreate) SaveX(ctx context.Context) *Blog {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlogCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlogCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlogCreate) check() error {
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Blog.title"`)}
	}
	if _, ok := bc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Blog.address"`)}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Blog.user_id"`)}
	}
	if _, ok := bc.mutation.Ctime(); !ok {
		return &ValidationError{Name: "ctime", err: errors.New(`ent: missing required field "Blog.ctime"`)}
	}
	if _, ok := bc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "Blog.desc"`)}
	}
	return nil
}

func (bc *BlogCreate) sqlSave(ctx context.Context) (*Blog, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlogCreate) createSpec() (*Blog, *sqlgraph.CreateSpec) {
	var (
		_node = &Blog{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blog.Table, sqlgraph.NewFieldSpec(blog.FieldID, field.TypeUint64))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Address(); ok {
		_spec.SetField(blog.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := bc.mutation.UserID(); ok {
		_spec.SetField(blog.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	if value, ok := bc.mutation.Ctime(); ok {
		_spec.SetField(blog.FieldCtime, field.TypeInt64, value)
		_node.Ctime = value
	}
	if value, ok := bc.mutation.Desc(); ok {
		_spec.SetField(blog.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	return _node, _spec
}

// BlogCreateBulk is the builder for creating many Blog entities in bulk.
type BlogCreateBulk struct {
	config
	err      error
	builders []*BlogCreate
}

// Save creates the Blog entities in the database.
func (bcb *BlogCreateBulk) Save(ctx context.Context) ([]*Blog, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blog, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlogMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlogCreateBulk) SaveX(ctx context.Context) []*Blog {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlogCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlogCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}