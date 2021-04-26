// Code generated by entc, DO NOT EDIT.

package ent

import (
	"boot/ent/tool"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ToolCreate is the builder for creating a Tool entity.
type ToolCreate struct {
	config
	mutation *ToolMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tc *ToolCreate) SetCreateTime(t time.Time) *ToolCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *ToolCreate) SetNillableCreateTime(t *time.Time) *ToolCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *ToolCreate) SetUpdateTime(t time.Time) *ToolCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *ToolCreate) SetNillableUpdateTime(t *time.Time) *ToolCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetMachineID sets the "machine_id" field.
func (tc *ToolCreate) SetMachineID(i int) *ToolCreate {
	tc.mutation.SetMachineID(i)
	return tc
}

// SetStatus sets the "status" field.
func (tc *ToolCreate) SetStatus(i int) *ToolCreate {
	tc.mutation.SetStatus(i)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *ToolCreate) SetNillableStatus(i *int) *ToolCreate {
	if i != nil {
		tc.SetStatus(*i)
	}
	return tc
}

// Mutation returns the ToolMutation object of the builder.
func (tc *ToolCreate) Mutation() *ToolMutation {
	return tc.mutation
}

// Save creates the Tool in the database.
func (tc *ToolCreate) Save(ctx context.Context) (*Tool, error) {
	var (
		err  error
		node *Tool
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ToolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ToolCreate) SaveX(ctx context.Context) *Tool {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tc *ToolCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := tool.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := tool.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := tool.DefaultStatus
		tc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ToolCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := tc.mutation.MachineID(); !ok {
		return &ValidationError{Name: "machine_id", err: errors.New("ent: missing required field \"machine_id\"")}
	}
	if v, ok := tc.mutation.MachineID(); ok {
		if err := tool.MachineIDValidator(v); err != nil {
			return &ValidationError{Name: "machine_id", err: fmt.Errorf("ent: validator failed for field \"machine_id\": %w", err)}
		}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	return nil
}

func (tc *ToolCreate) sqlSave(ctx context.Context) (*Tool, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *ToolCreate) createSpec() (*Tool, *sqlgraph.CreateSpec) {
	var (
		_node = &Tool{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tool.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tool.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tool.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.MachineID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tool.FieldMachineID,
		})
		_node.MachineID = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tool.FieldStatus,
		})
		_node.Status = value
	}
	return _node, _spec
}

// ToolCreateBulk is the builder for creating many Tool entities in bulk.
type ToolCreateBulk struct {
	config
	builders []*ToolCreate
}

// Save creates the Tool entities in the database.
func (tcb *ToolCreateBulk) Save(ctx context.Context) ([]*Tool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tool, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToolMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ToolCreateBulk) SaveX(ctx context.Context) []*Tool {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
