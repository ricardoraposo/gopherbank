// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/depositrequest"
	"github.com/ricardoraposo/gopherbank/ent/predicate"
)

// DepositRequestUpdate is the builder for updating DepositRequest entities.
type DepositRequestUpdate struct {
	config
	hooks    []Hook
	mutation *DepositRequestMutation
}

// Where appends a list predicates to the DepositRequestUpdate builder.
func (dru *DepositRequestUpdate) Where(ps ...predicate.DepositRequest) *DepositRequestUpdate {
	dru.mutation.Where(ps...)
	return dru
}

// SetAmount sets the "amount" field.
func (dru *DepositRequestUpdate) SetAmount(f float64) *DepositRequestUpdate {
	dru.mutation.ResetAmount()
	dru.mutation.SetAmount(f)
	return dru
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (dru *DepositRequestUpdate) SetNillableAmount(f *float64) *DepositRequestUpdate {
	if f != nil {
		dru.SetAmount(*f)
	}
	return dru
}

// AddAmount adds f to the "amount" field.
func (dru *DepositRequestUpdate) AddAmount(f float64) *DepositRequestUpdate {
	dru.mutation.AddAmount(f)
	return dru
}

// SetStatus sets the "status" field.
func (dru *DepositRequestUpdate) SetStatus(s string) *DepositRequestUpdate {
	dru.mutation.SetStatus(s)
	return dru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dru *DepositRequestUpdate) SetNillableStatus(s *string) *DepositRequestUpdate {
	if s != nil {
		dru.SetStatus(*s)
	}
	return dru
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (dru *DepositRequestUpdate) SetAccountID(id string) *DepositRequestUpdate {
	dru.mutation.SetAccountID(id)
	return dru
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (dru *DepositRequestUpdate) SetNillableAccountID(id *string) *DepositRequestUpdate {
	if id != nil {
		dru = dru.SetAccountID(*id)
	}
	return dru
}

// SetAccount sets the "account" edge to the Account entity.
func (dru *DepositRequestUpdate) SetAccount(a *Account) *DepositRequestUpdate {
	return dru.SetAccountID(a.ID)
}

// Mutation returns the DepositRequestMutation object of the builder.
func (dru *DepositRequestUpdate) Mutation() *DepositRequestMutation {
	return dru.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (dru *DepositRequestUpdate) ClearAccount() *DepositRequestUpdate {
	dru.mutation.ClearAccount()
	return dru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dru *DepositRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dru.sqlSave, dru.mutation, dru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dru *DepositRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := dru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dru *DepositRequestUpdate) Exec(ctx context.Context) error {
	_, err := dru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dru *DepositRequestUpdate) ExecX(ctx context.Context) {
	if err := dru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dru *DepositRequestUpdate) check() error {
	if v, ok := dru.mutation.Amount(); ok {
		if err := depositrequest.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "DepositRequest.amount": %w`, err)}
		}
	}
	if v, ok := dru.mutation.Status(); ok {
		if err := depositrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "DepositRequest.status": %w`, err)}
		}
	}
	return nil
}

func (dru *DepositRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(depositrequest.Table, depositrequest.Columns, sqlgraph.NewFieldSpec(depositrequest.FieldID, field.TypeInt))
	if ps := dru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dru.mutation.Amount(); ok {
		_spec.SetField(depositrequest.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := dru.mutation.AddedAmount(); ok {
		_spec.AddField(depositrequest.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := dru.mutation.Status(); ok {
		_spec.SetField(depositrequest.FieldStatus, field.TypeString, value)
	}
	if dru.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   depositrequest.AccountTable,
			Columns: []string{depositrequest.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   depositrequest.AccountTable,
			Columns: []string{depositrequest.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{depositrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dru.mutation.done = true
	return n, nil
}

// DepositRequestUpdateOne is the builder for updating a single DepositRequest entity.
type DepositRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DepositRequestMutation
}

// SetAmount sets the "amount" field.
func (druo *DepositRequestUpdateOne) SetAmount(f float64) *DepositRequestUpdateOne {
	druo.mutation.ResetAmount()
	druo.mutation.SetAmount(f)
	return druo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (druo *DepositRequestUpdateOne) SetNillableAmount(f *float64) *DepositRequestUpdateOne {
	if f != nil {
		druo.SetAmount(*f)
	}
	return druo
}

// AddAmount adds f to the "amount" field.
func (druo *DepositRequestUpdateOne) AddAmount(f float64) *DepositRequestUpdateOne {
	druo.mutation.AddAmount(f)
	return druo
}

// SetStatus sets the "status" field.
func (druo *DepositRequestUpdateOne) SetStatus(s string) *DepositRequestUpdateOne {
	druo.mutation.SetStatus(s)
	return druo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (druo *DepositRequestUpdateOne) SetNillableStatus(s *string) *DepositRequestUpdateOne {
	if s != nil {
		druo.SetStatus(*s)
	}
	return druo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (druo *DepositRequestUpdateOne) SetAccountID(id string) *DepositRequestUpdateOne {
	druo.mutation.SetAccountID(id)
	return druo
}

// SetNillableAccountID sets the "account" edge to the Account entity by ID if the given value is not nil.
func (druo *DepositRequestUpdateOne) SetNillableAccountID(id *string) *DepositRequestUpdateOne {
	if id != nil {
		druo = druo.SetAccountID(*id)
	}
	return druo
}

// SetAccount sets the "account" edge to the Account entity.
func (druo *DepositRequestUpdateOne) SetAccount(a *Account) *DepositRequestUpdateOne {
	return druo.SetAccountID(a.ID)
}

// Mutation returns the DepositRequestMutation object of the builder.
func (druo *DepositRequestUpdateOne) Mutation() *DepositRequestMutation {
	return druo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (druo *DepositRequestUpdateOne) ClearAccount() *DepositRequestUpdateOne {
	druo.mutation.ClearAccount()
	return druo
}

// Where appends a list predicates to the DepositRequestUpdate builder.
func (druo *DepositRequestUpdateOne) Where(ps ...predicate.DepositRequest) *DepositRequestUpdateOne {
	druo.mutation.Where(ps...)
	return druo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (druo *DepositRequestUpdateOne) Select(field string, fields ...string) *DepositRequestUpdateOne {
	druo.fields = append([]string{field}, fields...)
	return druo
}

// Save executes the query and returns the updated DepositRequest entity.
func (druo *DepositRequestUpdateOne) Save(ctx context.Context) (*DepositRequest, error) {
	return withHooks(ctx, druo.sqlSave, druo.mutation, druo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (druo *DepositRequestUpdateOne) SaveX(ctx context.Context) *DepositRequest {
	node, err := druo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (druo *DepositRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := druo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (druo *DepositRequestUpdateOne) ExecX(ctx context.Context) {
	if err := druo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (druo *DepositRequestUpdateOne) check() error {
	if v, ok := druo.mutation.Amount(); ok {
		if err := depositrequest.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "DepositRequest.amount": %w`, err)}
		}
	}
	if v, ok := druo.mutation.Status(); ok {
		if err := depositrequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "DepositRequest.status": %w`, err)}
		}
	}
	return nil
}

func (druo *DepositRequestUpdateOne) sqlSave(ctx context.Context) (_node *DepositRequest, err error) {
	if err := druo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(depositrequest.Table, depositrequest.Columns, sqlgraph.NewFieldSpec(depositrequest.FieldID, field.TypeInt))
	id, ok := druo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DepositRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := druo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, depositrequest.FieldID)
		for _, f := range fields {
			if !depositrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != depositrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := druo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := druo.mutation.Amount(); ok {
		_spec.SetField(depositrequest.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := druo.mutation.AddedAmount(); ok {
		_spec.AddField(depositrequest.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := druo.mutation.Status(); ok {
		_spec.SetField(depositrequest.FieldStatus, field.TypeString, value)
	}
	if druo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   depositrequest.AccountTable,
			Columns: []string{depositrequest.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   depositrequest.AccountTable,
			Columns: []string{depositrequest.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DepositRequest{config: druo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, druo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{depositrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	druo.mutation.done = true
	return _node, nil
}