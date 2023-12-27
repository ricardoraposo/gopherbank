// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ricardoraposo/gopherbank/ent/predicate"
	"github.com/ricardoraposo/gopherbank/ent/transactiondetails"
)

// TransactionDetailsDelete is the builder for deleting a TransactionDetails entity.
type TransactionDetailsDelete struct {
	config
	hooks    []Hook
	mutation *TransactionDetailsMutation
}

// Where appends a list predicates to the TransactionDetailsDelete builder.
func (tdd *TransactionDetailsDelete) Where(ps ...predicate.TransactionDetails) *TransactionDetailsDelete {
	tdd.mutation.Where(ps...)
	return tdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tdd *TransactionDetailsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tdd.sqlExec, tdd.mutation, tdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tdd *TransactionDetailsDelete) ExecX(ctx context.Context) int {
	n, err := tdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tdd *TransactionDetailsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(transactiondetails.Table, sqlgraph.NewFieldSpec(transactiondetails.FieldID, field.TypeInt))
	if ps := tdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tdd.mutation.done = true
	return affected, err
}

// TransactionDetailsDeleteOne is the builder for deleting a single TransactionDetails entity.
type TransactionDetailsDeleteOne struct {
	tdd *TransactionDetailsDelete
}

// Where appends a list predicates to the TransactionDetailsDelete builder.
func (tddo *TransactionDetailsDeleteOne) Where(ps ...predicate.TransactionDetails) *TransactionDetailsDeleteOne {
	tddo.tdd.mutation.Where(ps...)
	return tddo
}

// Exec executes the deletion query.
func (tddo *TransactionDetailsDeleteOne) Exec(ctx context.Context) error {
	n, err := tddo.tdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{transactiondetails.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tddo *TransactionDetailsDeleteOne) ExecX(ctx context.Context) {
	if err := tddo.Exec(ctx); err != nil {
		panic(err)
	}
}