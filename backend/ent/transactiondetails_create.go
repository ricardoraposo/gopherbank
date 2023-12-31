// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ricardoraposo/gopherbank/ent/transaction"
	"github.com/ricardoraposo/gopherbank/ent/transactiondetails"
)

// TransactionDetailsCreate is the builder for creating a TransactionDetails entity.
type TransactionDetailsCreate struct {
	config
	mutation *TransactionDetailsMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (tdc *TransactionDetailsCreate) SetAmount(f float64) *TransactionDetailsCreate {
	tdc.mutation.SetAmount(f)
	return tdc
}

// SetType sets the "type" field.
func (tdc *TransactionDetailsCreate) SetType(s string) *TransactionDetailsCreate {
	tdc.mutation.SetType(s)
	return tdc
}

// SetTransactedAt sets the "transacted_at" field.
func (tdc *TransactionDetailsCreate) SetTransactedAt(t time.Time) *TransactionDetailsCreate {
	tdc.mutation.SetTransactedAt(t)
	return tdc
}

// SetNillableTransactedAt sets the "transacted_at" field if the given value is not nil.
func (tdc *TransactionDetailsCreate) SetNillableTransactedAt(t *time.Time) *TransactionDetailsCreate {
	if t != nil {
		tdc.SetTransactedAt(*t)
	}
	return tdc
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (tdc *TransactionDetailsCreate) SetTransactionID(id int) *TransactionDetailsCreate {
	tdc.mutation.SetTransactionID(id)
	return tdc
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (tdc *TransactionDetailsCreate) SetTransaction(t *Transaction) *TransactionDetailsCreate {
	return tdc.SetTransactionID(t.ID)
}

// Mutation returns the TransactionDetailsMutation object of the builder.
func (tdc *TransactionDetailsCreate) Mutation() *TransactionDetailsMutation {
	return tdc.mutation
}

// Save creates the TransactionDetails in the database.
func (tdc *TransactionDetailsCreate) Save(ctx context.Context) (*TransactionDetails, error) {
	tdc.defaults()
	return withHooks(ctx, tdc.sqlSave, tdc.mutation, tdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tdc *TransactionDetailsCreate) SaveX(ctx context.Context) *TransactionDetails {
	v, err := tdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdc *TransactionDetailsCreate) Exec(ctx context.Context) error {
	_, err := tdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdc *TransactionDetailsCreate) ExecX(ctx context.Context) {
	if err := tdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdc *TransactionDetailsCreate) defaults() {
	if _, ok := tdc.mutation.TransactedAt(); !ok {
		v := transactiondetails.DefaultTransactedAt()
		tdc.mutation.SetTransactedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdc *TransactionDetailsCreate) check() error {
	if _, ok := tdc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "TransactionDetails.amount"`)}
	}
	if _, ok := tdc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "TransactionDetails.type"`)}
	}
	if v, ok := tdc.mutation.GetType(); ok {
		if err := transactiondetails.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "TransactionDetails.type": %w`, err)}
		}
	}
	if _, ok := tdc.mutation.TransactedAt(); !ok {
		return &ValidationError{Name: "transacted_at", err: errors.New(`ent: missing required field "TransactionDetails.transacted_at"`)}
	}
	if _, ok := tdc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction", err: errors.New(`ent: missing required edge "TransactionDetails.transaction"`)}
	}
	return nil
}

func (tdc *TransactionDetailsCreate) sqlSave(ctx context.Context) (*TransactionDetails, error) {
	if err := tdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tdc.mutation.id = &_node.ID
	tdc.mutation.done = true
	return _node, nil
}

func (tdc *TransactionDetailsCreate) createSpec() (*TransactionDetails, *sqlgraph.CreateSpec) {
	var (
		_node = &TransactionDetails{config: tdc.config}
		_spec = sqlgraph.NewCreateSpec(transactiondetails.Table, sqlgraph.NewFieldSpec(transactiondetails.FieldID, field.TypeInt))
	)
	if value, ok := tdc.mutation.Amount(); ok {
		_spec.SetField(transactiondetails.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := tdc.mutation.GetType(); ok {
		_spec.SetField(transactiondetails.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := tdc.mutation.TransactedAt(); ok {
		_spec.SetField(transactiondetails.FieldTransactedAt, field.TypeTime, value)
		_node.TransactedAt = value
	}
	if nodes := tdc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transactiondetails.TransactionTable,
			Columns: []string{transactiondetails.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.transaction_detail = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionDetailsCreateBulk is the builder for creating many TransactionDetails entities in bulk.
type TransactionDetailsCreateBulk struct {
	config
	err      error
	builders []*TransactionDetailsCreate
}

// Save creates the TransactionDetails entities in the database.
func (tdcb *TransactionDetailsCreateBulk) Save(ctx context.Context) ([]*TransactionDetails, error) {
	if tdcb.err != nil {
		return nil, tdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tdcb.builders))
	nodes := make([]*TransactionDetails, len(tdcb.builders))
	mutators := make([]Mutator, len(tdcb.builders))
	for i := range tdcb.builders {
		func(i int, root context.Context) {
			builder := tdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionDetailsMutation)
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
					_, err = mutators[i+1].Mutate(root, tdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, tdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tdcb *TransactionDetailsCreateBulk) SaveX(ctx context.Context) []*TransactionDetails {
	v, err := tdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdcb *TransactionDetailsCreateBulk) Exec(ctx context.Context) error {
	_, err := tdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdcb *TransactionDetailsCreateBulk) ExecX(ctx context.Context) {
	if err := tdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
