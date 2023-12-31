// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ricardoraposo/gopherbank/ent/predicate"
	"github.com/ricardoraposo/gopherbank/ent/transaction"
	"github.com/ricardoraposo/gopherbank/ent/transactiondetails"
)

// TransactionDetailsUpdate is the builder for updating TransactionDetails entities.
type TransactionDetailsUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionDetailsMutation
}

// Where appends a list predicates to the TransactionDetailsUpdate builder.
func (tdu *TransactionDetailsUpdate) Where(ps ...predicate.TransactionDetails) *TransactionDetailsUpdate {
	tdu.mutation.Where(ps...)
	return tdu
}

// SetAmount sets the "amount" field.
func (tdu *TransactionDetailsUpdate) SetAmount(f float64) *TransactionDetailsUpdate {
	tdu.mutation.ResetAmount()
	tdu.mutation.SetAmount(f)
	return tdu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tdu *TransactionDetailsUpdate) SetNillableAmount(f *float64) *TransactionDetailsUpdate {
	if f != nil {
		tdu.SetAmount(*f)
	}
	return tdu
}

// AddAmount adds f to the "amount" field.
func (tdu *TransactionDetailsUpdate) AddAmount(f float64) *TransactionDetailsUpdate {
	tdu.mutation.AddAmount(f)
	return tdu
}

// SetType sets the "type" field.
func (tdu *TransactionDetailsUpdate) SetType(s string) *TransactionDetailsUpdate {
	tdu.mutation.SetType(s)
	return tdu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tdu *TransactionDetailsUpdate) SetNillableType(s *string) *TransactionDetailsUpdate {
	if s != nil {
		tdu.SetType(*s)
	}
	return tdu
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (tdu *TransactionDetailsUpdate) SetTransactionID(id int) *TransactionDetailsUpdate {
	tdu.mutation.SetTransactionID(id)
	return tdu
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (tdu *TransactionDetailsUpdate) SetTransaction(t *Transaction) *TransactionDetailsUpdate {
	return tdu.SetTransactionID(t.ID)
}

// Mutation returns the TransactionDetailsMutation object of the builder.
func (tdu *TransactionDetailsUpdate) Mutation() *TransactionDetailsMutation {
	return tdu.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (tdu *TransactionDetailsUpdate) ClearTransaction() *TransactionDetailsUpdate {
	tdu.mutation.ClearTransaction()
	return tdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tdu *TransactionDetailsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tdu.sqlSave, tdu.mutation, tdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tdu *TransactionDetailsUpdate) SaveX(ctx context.Context) int {
	affected, err := tdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tdu *TransactionDetailsUpdate) Exec(ctx context.Context) error {
	_, err := tdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdu *TransactionDetailsUpdate) ExecX(ctx context.Context) {
	if err := tdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdu *TransactionDetailsUpdate) check() error {
	if v, ok := tdu.mutation.GetType(); ok {
		if err := transactiondetails.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "TransactionDetails.type": %w`, err)}
		}
	}
	if _, ok := tdu.mutation.TransactionID(); tdu.mutation.TransactionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TransactionDetails.transaction"`)
	}
	return nil
}

func (tdu *TransactionDetailsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(transactiondetails.Table, transactiondetails.Columns, sqlgraph.NewFieldSpec(transactiondetails.FieldID, field.TypeInt))
	if ps := tdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdu.mutation.Amount(); ok {
		_spec.SetField(transactiondetails.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedAmount(); ok {
		_spec.AddField(transactiondetails.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.GetType(); ok {
		_spec.SetField(transactiondetails.FieldType, field.TypeString, value)
	}
	if tdu.mutation.TransactionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tdu.mutation.TransactionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiondetails.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tdu.mutation.done = true
	return n, nil
}

// TransactionDetailsUpdateOne is the builder for updating a single TransactionDetails entity.
type TransactionDetailsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionDetailsMutation
}

// SetAmount sets the "amount" field.
func (tduo *TransactionDetailsUpdateOne) SetAmount(f float64) *TransactionDetailsUpdateOne {
	tduo.mutation.ResetAmount()
	tduo.mutation.SetAmount(f)
	return tduo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tduo *TransactionDetailsUpdateOne) SetNillableAmount(f *float64) *TransactionDetailsUpdateOne {
	if f != nil {
		tduo.SetAmount(*f)
	}
	return tduo
}

// AddAmount adds f to the "amount" field.
func (tduo *TransactionDetailsUpdateOne) AddAmount(f float64) *TransactionDetailsUpdateOne {
	tduo.mutation.AddAmount(f)
	return tduo
}

// SetType sets the "type" field.
func (tduo *TransactionDetailsUpdateOne) SetType(s string) *TransactionDetailsUpdateOne {
	tduo.mutation.SetType(s)
	return tduo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tduo *TransactionDetailsUpdateOne) SetNillableType(s *string) *TransactionDetailsUpdateOne {
	if s != nil {
		tduo.SetType(*s)
	}
	return tduo
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (tduo *TransactionDetailsUpdateOne) SetTransactionID(id int) *TransactionDetailsUpdateOne {
	tduo.mutation.SetTransactionID(id)
	return tduo
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (tduo *TransactionDetailsUpdateOne) SetTransaction(t *Transaction) *TransactionDetailsUpdateOne {
	return tduo.SetTransactionID(t.ID)
}

// Mutation returns the TransactionDetailsMutation object of the builder.
func (tduo *TransactionDetailsUpdateOne) Mutation() *TransactionDetailsMutation {
	return tduo.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (tduo *TransactionDetailsUpdateOne) ClearTransaction() *TransactionDetailsUpdateOne {
	tduo.mutation.ClearTransaction()
	return tduo
}

// Where appends a list predicates to the TransactionDetailsUpdate builder.
func (tduo *TransactionDetailsUpdateOne) Where(ps ...predicate.TransactionDetails) *TransactionDetailsUpdateOne {
	tduo.mutation.Where(ps...)
	return tduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tduo *TransactionDetailsUpdateOne) Select(field string, fields ...string) *TransactionDetailsUpdateOne {
	tduo.fields = append([]string{field}, fields...)
	return tduo
}

// Save executes the query and returns the updated TransactionDetails entity.
func (tduo *TransactionDetailsUpdateOne) Save(ctx context.Context) (*TransactionDetails, error) {
	return withHooks(ctx, tduo.sqlSave, tduo.mutation, tduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tduo *TransactionDetailsUpdateOne) SaveX(ctx context.Context) *TransactionDetails {
	node, err := tduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tduo *TransactionDetailsUpdateOne) Exec(ctx context.Context) error {
	_, err := tduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tduo *TransactionDetailsUpdateOne) ExecX(ctx context.Context) {
	if err := tduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tduo *TransactionDetailsUpdateOne) check() error {
	if v, ok := tduo.mutation.GetType(); ok {
		if err := transactiondetails.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "TransactionDetails.type": %w`, err)}
		}
	}
	if _, ok := tduo.mutation.TransactionID(); tduo.mutation.TransactionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TransactionDetails.transaction"`)
	}
	return nil
}

func (tduo *TransactionDetailsUpdateOne) sqlSave(ctx context.Context) (_node *TransactionDetails, err error) {
	if err := tduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(transactiondetails.Table, transactiondetails.Columns, sqlgraph.NewFieldSpec(transactiondetails.FieldID, field.TypeInt))
	id, ok := tduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TransactionDetails.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactiondetails.FieldID)
		for _, f := range fields {
			if !transactiondetails.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transactiondetails.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tduo.mutation.Amount(); ok {
		_spec.SetField(transactiondetails.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedAmount(); ok {
		_spec.AddField(transactiondetails.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.GetType(); ok {
		_spec.SetField(transactiondetails.FieldType, field.TypeString, value)
	}
	if tduo.mutation.TransactionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tduo.mutation.TransactionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TransactionDetails{config: tduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiondetails.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tduo.mutation.done = true
	return _node, nil
}
