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
	"github.com/ricardoraposo/gopherbank/ent/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePassword(s *string) *AccountUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// SetBalance sets the "balance" field.
func (au *AccountUpdate) SetBalance(f float64) *AccountUpdate {
	au.mutation.ResetBalance()
	au.mutation.SetBalance(f)
	return au
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (au *AccountUpdate) SetNillableBalance(f *float64) *AccountUpdate {
	if f != nil {
		au.SetBalance(*f)
	}
	return au
}

// AddBalance adds f to the "balance" field.
func (au *AccountUpdate) AddBalance(f float64) *AccountUpdate {
	au.mutation.AddBalance(f)
	return au
}

// SetAdmin sets the "admin" field.
func (au *AccountUpdate) SetAdmin(b bool) *AccountUpdate {
	au.mutation.SetAdmin(b)
	return au
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAdmin(b *bool) *AccountUpdate {
	if b != nil {
		au.SetAdmin(*b)
	}
	return au
}

// AddFavoritedIDs adds the "favoriteds" edge to the Account entity by IDs.
func (au *AccountUpdate) AddFavoritedIDs(ids ...string) *AccountUpdate {
	au.mutation.AddFavoritedIDs(ids...)
	return au
}

// AddFavoriteds adds the "favoriteds" edges to the Account entity.
func (au *AccountUpdate) AddFavoriteds(a ...*Account) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddFavoritedIDs(ids...)
}

// AddFavoriteIDs adds the "favorites" edge to the Account entity by IDs.
func (au *AccountUpdate) AddFavoriteIDs(ids ...string) *AccountUpdate {
	au.mutation.AddFavoriteIDs(ids...)
	return au
}

// AddFavorites adds the "favorites" edges to the Account entity.
func (au *AccountUpdate) AddFavorites(a ...*Account) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddFavoriteIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearFavoriteds clears all "favoriteds" edges to the Account entity.
func (au *AccountUpdate) ClearFavoriteds() *AccountUpdate {
	au.mutation.ClearFavoriteds()
	return au
}

// RemoveFavoritedIDs removes the "favoriteds" edge to Account entities by IDs.
func (au *AccountUpdate) RemoveFavoritedIDs(ids ...string) *AccountUpdate {
	au.mutation.RemoveFavoritedIDs(ids...)
	return au
}

// RemoveFavoriteds removes "favoriteds" edges to Account entities.
func (au *AccountUpdate) RemoveFavoriteds(a ...*Account) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveFavoritedIDs(ids...)
}

// ClearFavorites clears all "favorites" edges to the Account entity.
func (au *AccountUpdate) ClearFavorites() *AccountUpdate {
	au.mutation.ClearFavorites()
	return au
}

// RemoveFavoriteIDs removes the "favorites" edge to Account entities by IDs.
func (au *AccountUpdate) RemoveFavoriteIDs(ids ...string) *AccountUpdate {
	au.mutation.RemoveFavoriteIDs(ids...)
	return au
}

// RemoveFavorites removes "favorites" edges to Account entities.
func (au *AccountUpdate) RemoveFavorites(a ...*Account) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveFavoriteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.Balance(); ok {
		_spec.SetField(account.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedBalance(); ok {
		_spec.AddField(account.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.Admin(); ok {
		_spec.SetField(account.FieldAdmin, field.TypeBool, value)
	}
	if au.mutation.FavoritedsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.FavoritedsTable,
			Columns: account.FavoritedsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedFavoritedsIDs(); len(nodes) > 0 && !au.mutation.FavoritedsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.FavoritedsTable,
			Columns: account.FavoritedsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.FavoritedsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.FavoritedsTable,
			Columns: account.FavoritedsPrimaryKey,
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
	if au.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FavoritesTable,
			Columns: account.FavoritesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedFavoritesIDs(); len(nodes) > 0 && !au.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FavoritesTable,
			Columns: account.FavoritesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FavoritesTable,
			Columns: account.FavoritesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePassword(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// SetBalance sets the "balance" field.
func (auo *AccountUpdateOne) SetBalance(f float64) *AccountUpdateOne {
	auo.mutation.ResetBalance()
	auo.mutation.SetBalance(f)
	return auo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableBalance(f *float64) *AccountUpdateOne {
	if f != nil {
		auo.SetBalance(*f)
	}
	return auo
}

// AddBalance adds f to the "balance" field.
func (auo *AccountUpdateOne) AddBalance(f float64) *AccountUpdateOne {
	auo.mutation.AddBalance(f)
	return auo
}

// SetAdmin sets the "admin" field.
func (auo *AccountUpdateOne) SetAdmin(b bool) *AccountUpdateOne {
	auo.mutation.SetAdmin(b)
	return auo
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAdmin(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetAdmin(*b)
	}
	return auo
}

// AddFavoritedIDs adds the "favoriteds" edge to the Account entity by IDs.
func (auo *AccountUpdateOne) AddFavoritedIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.AddFavoritedIDs(ids...)
	return auo
}

// AddFavoriteds adds the "favoriteds" edges to the Account entity.
func (auo *AccountUpdateOne) AddFavoriteds(a ...*Account) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddFavoritedIDs(ids...)
}

// AddFavoriteIDs adds the "favorites" edge to the Account entity by IDs.
func (auo *AccountUpdateOne) AddFavoriteIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.AddFavoriteIDs(ids...)
	return auo
}

// AddFavorites adds the "favorites" edges to the Account entity.
func (auo *AccountUpdateOne) AddFavorites(a ...*Account) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddFavoriteIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearFavoriteds clears all "favoriteds" edges to the Account entity.
func (auo *AccountUpdateOne) ClearFavoriteds() *AccountUpdateOne {
	auo.mutation.ClearFavoriteds()
	return auo
}

// RemoveFavoritedIDs removes the "favoriteds" edge to Account entities by IDs.
func (auo *AccountUpdateOne) RemoveFavoritedIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.RemoveFavoritedIDs(ids...)
	return auo
}

// RemoveFavoriteds removes "favoriteds" edges to Account entities.
func (auo *AccountUpdateOne) RemoveFavoriteds(a ...*Account) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveFavoritedIDs(ids...)
}

// ClearFavorites clears all "favorites" edges to the Account entity.
func (auo *AccountUpdateOne) ClearFavorites() *AccountUpdateOne {
	auo.mutation.ClearFavorites()
	return auo
}

// RemoveFavoriteIDs removes the "favorites" edge to Account entities by IDs.
func (auo *AccountUpdateOne) RemoveFavoriteIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.RemoveFavoriteIDs(ids...)
	return auo
}

// RemoveFavorites removes "favorites" edges to Account entities.
func (auo *AccountUpdateOne) RemoveFavorites(a ...*Account) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveFavoriteIDs(ids...)
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.Balance(); ok {
		_spec.SetField(account.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedBalance(); ok {
		_spec.AddField(account.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.Admin(); ok {
		_spec.SetField(account.FieldAdmin, field.TypeBool, value)
	}
	if auo.mutation.FavoritedsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.FavoritedsTable,
			Columns: account.FavoritedsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedFavoritedsIDs(); len(nodes) > 0 && !auo.mutation.FavoritedsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.FavoritedsTable,
			Columns: account.FavoritedsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.FavoritedsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.FavoritedsTable,
			Columns: account.FavoritedsPrimaryKey,
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
	if auo.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FavoritesTable,
			Columns: account.FavoritesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedFavoritesIDs(); len(nodes) > 0 && !auo.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FavoritesTable,
			Columns: account.FavoritesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FavoritesTable,
			Columns: account.FavoritesPrimaryKey,
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
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
