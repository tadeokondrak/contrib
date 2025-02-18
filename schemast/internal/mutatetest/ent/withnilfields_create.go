// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/withnilfields"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithNilFieldsCreate is the builder for creating a WithNilFields entity.
type WithNilFieldsCreate struct {
	config
	mutation *WithNilFieldsMutation
	hooks    []Hook
}

// Mutation returns the WithNilFieldsMutation object of the builder.
func (wnfc *WithNilFieldsCreate) Mutation() *WithNilFieldsMutation {
	return wnfc.mutation
}

// Save creates the WithNilFields in the database.
func (wnfc *WithNilFieldsCreate) Save(ctx context.Context) (*WithNilFields, error) {
	var (
		err  error
		node *WithNilFields
	)
	if len(wnfc.hooks) == 0 {
		if err = wnfc.check(); err != nil {
			return nil, err
		}
		node, err = wnfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithNilFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wnfc.check(); err != nil {
				return nil, err
			}
			wnfc.mutation = mutation
			if node, err = wnfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wnfc.hooks) - 1; i >= 0; i-- {
			if wnfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wnfc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*WithNilFields)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WithNilFieldsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wnfc *WithNilFieldsCreate) SaveX(ctx context.Context) *WithNilFields {
	v, err := wnfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wnfc *WithNilFieldsCreate) Exec(ctx context.Context) error {
	_, err := wnfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnfc *WithNilFieldsCreate) ExecX(ctx context.Context) {
	if err := wnfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wnfc *WithNilFieldsCreate) check() error {
	return nil
}

func (wnfc *WithNilFieldsCreate) sqlSave(ctx context.Context) (*WithNilFields, error) {
	_node, _spec := wnfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wnfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wnfc *WithNilFieldsCreate) createSpec() (*WithNilFields, *sqlgraph.CreateSpec) {
	var (
		_node = &WithNilFields{config: wnfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: withnilfields.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withnilfields.FieldID,
			},
		}
	)
	return _node, _spec
}

// WithNilFieldsCreateBulk is the builder for creating many WithNilFields entities in bulk.
type WithNilFieldsCreateBulk struct {
	config
	builders []*WithNilFieldsCreate
}

// Save creates the WithNilFields entities in the database.
func (wnfcb *WithNilFieldsCreateBulk) Save(ctx context.Context) ([]*WithNilFields, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wnfcb.builders))
	nodes := make([]*WithNilFields, len(wnfcb.builders))
	mutators := make([]Mutator, len(wnfcb.builders))
	for i := range wnfcb.builders {
		func(i int, root context.Context) {
			builder := wnfcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WithNilFieldsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wnfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wnfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wnfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wnfcb *WithNilFieldsCreateBulk) SaveX(ctx context.Context) []*WithNilFields {
	v, err := wnfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wnfcb *WithNilFieldsCreateBulk) Exec(ctx context.Context) error {
	_, err := wnfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnfcb *WithNilFieldsCreateBulk) ExecX(ctx context.Context) {
	if err := wnfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
