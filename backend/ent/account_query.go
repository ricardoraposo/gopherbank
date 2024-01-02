// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/depositrequest"
	"github.com/ricardoraposo/gopherbank/ent/notification"
	"github.com/ricardoraposo/gopherbank/ent/predicate"
	"github.com/ricardoraposo/gopherbank/ent/transaction"
	"github.com/ricardoraposo/gopherbank/ent/user"
)

// AccountQuery is the builder for querying Account entities.
type AccountQuery struct {
	config
	ctx                *QueryContext
	order              []account.OrderOption
	inters             []Interceptor
	predicates         []predicate.Account
	withUser           *UserQuery
	withFavoriteds     *AccountQuery
	withFavorites      *AccountQuery
	withFromAccount    *TransactionQuery
	withToAccount      *TransactionQuery
	withDepositRequest *DepositRequestQuery
	withNotification   *NotificationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountQuery builder.
func (aq *AccountQuery) Where(ps ...predicate.Account) *AccountQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AccountQuery) Limit(limit int) *AccountQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AccountQuery) Offset(offset int) *AccountQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AccountQuery) Unique(unique bool) *AccountQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AccountQuery) Order(o ...account.OrderOption) *AccountQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *AccountQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, account.UserTable, account.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFavoriteds chains the current query on the "favoriteds" edge.
func (aq *AccountQuery) QueryFavoriteds() *AccountQuery {
	query := (&AccountClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, account.FavoritedsTable, account.FavoritedsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFavorites chains the current query on the "favorites" edge.
func (aq *AccountQuery) QueryFavorites() *AccountQuery {
	query := (&AccountClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, account.FavoritesTable, account.FavoritesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFromAccount chains the current query on the "from_account" edge.
func (aq *AccountQuery) QueryFromAccount() *TransactionQuery {
	query := (&TransactionClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.FromAccountTable, account.FromAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToAccount chains the current query on the "to_account" edge.
func (aq *AccountQuery) QueryToAccount() *TransactionQuery {
	query := (&TransactionClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.ToAccountTable, account.ToAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepositRequest chains the current query on the "deposit_request" edge.
func (aq *AccountQuery) QueryDepositRequest() *DepositRequestQuery {
	query := (&DepositRequestClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(depositrequest.Table, depositrequest.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.DepositRequestTable, account.DepositRequestColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotification chains the current query on the "notification" edge.
func (aq *AccountQuery) QueryNotification() *NotificationQuery {
	query := (&NotificationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.NotificationTable, account.NotificationColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Account entity from the query.
// Returns a *NotFoundError when no Account was found.
func (aq *AccountQuery) First(ctx context.Context) (*Account, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{account.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AccountQuery) FirstX(ctx context.Context) *Account {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Account ID from the query.
// Returns a *NotFoundError when no Account ID was found.
func (aq *AccountQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{account.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AccountQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Account entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Account entity is found.
// Returns a *NotFoundError when no Account entities are found.
func (aq *AccountQuery) Only(ctx context.Context) (*Account, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{account.Label}
	default:
		return nil, &NotSingularError{account.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AccountQuery) OnlyX(ctx context.Context) *Account {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Account ID in the query.
// Returns a *NotSingularError when more than one Account ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AccountQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{account.Label}
	default:
		err = &NotSingularError{account.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AccountQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Accounts.
func (aq *AccountQuery) All(ctx context.Context) ([]*Account, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Account, *AccountQuery]()
	return withInterceptors[[]*Account](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AccountQuery) AllX(ctx context.Context) []*Account {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Account IDs.
func (aq *AccountQuery) IDs(ctx context.Context) (ids []string, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(account.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AccountQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AccountQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AccountQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AccountQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AccountQuery) Clone() *AccountQuery {
	if aq == nil {
		return nil
	}
	return &AccountQuery{
		config:             aq.config,
		ctx:                aq.ctx.Clone(),
		order:              append([]account.OrderOption{}, aq.order...),
		inters:             append([]Interceptor{}, aq.inters...),
		predicates:         append([]predicate.Account{}, aq.predicates...),
		withUser:           aq.withUser.Clone(),
		withFavoriteds:     aq.withFavoriteds.Clone(),
		withFavorites:      aq.withFavorites.Clone(),
		withFromAccount:    aq.withFromAccount.Clone(),
		withToAccount:      aq.withToAccount.Clone(),
		withDepositRequest: aq.withDepositRequest.Clone(),
		withNotification:   aq.withNotification.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithUser(opts ...func(*UserQuery)) *AccountQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// WithFavoriteds tells the query-builder to eager-load the nodes that are connected to
// the "favoriteds" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithFavoriteds(opts ...func(*AccountQuery)) *AccountQuery {
	query := (&AccountClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withFavoriteds = query
	return aq
}

// WithFavorites tells the query-builder to eager-load the nodes that are connected to
// the "favorites" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithFavorites(opts ...func(*AccountQuery)) *AccountQuery {
	query := (&AccountClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withFavorites = query
	return aq
}

// WithFromAccount tells the query-builder to eager-load the nodes that are connected to
// the "from_account" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithFromAccount(opts ...func(*TransactionQuery)) *AccountQuery {
	query := (&TransactionClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withFromAccount = query
	return aq
}

// WithToAccount tells the query-builder to eager-load the nodes that are connected to
// the "to_account" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithToAccount(opts ...func(*TransactionQuery)) *AccountQuery {
	query := (&TransactionClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withToAccount = query
	return aq
}

// WithDepositRequest tells the query-builder to eager-load the nodes that are connected to
// the "deposit_request" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithDepositRequest(opts ...func(*DepositRequestQuery)) *AccountQuery {
	query := (&DepositRequestClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withDepositRequest = query
	return aq
}

// WithNotification tells the query-builder to eager-load the nodes that are connected to
// the "notification" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithNotification(opts ...func(*NotificationQuery)) *AccountQuery {
	query := (&NotificationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withNotification = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Password string `json:"-"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Account.Query().
//		GroupBy(account.FieldPassword).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AccountQuery) GroupBy(field string, fields ...string) *AccountGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = account.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Password string `json:"-"`
//	}
//
//	client.Account.Query().
//		Select(account.FieldPassword).
//		Scan(ctx, &v)
func (aq *AccountQuery) Select(fields ...string) *AccountSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AccountSelect{AccountQuery: aq}
	sbuild.label = account.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountSelect configured with the given aggregations.
func (aq *AccountQuery) Aggregate(fns ...AggregateFunc) *AccountSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !account.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Account, error) {
	var (
		nodes       = []*Account{}
		_spec       = aq.querySpec()
		loadedTypes = [7]bool{
			aq.withUser != nil,
			aq.withFavoriteds != nil,
			aq.withFavorites != nil,
			aq.withFromAccount != nil,
			aq.withToAccount != nil,
			aq.withDepositRequest != nil,
			aq.withNotification != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Account).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Account{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes, nil,
			func(n *Account, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withFavoriteds; query != nil {
		if err := aq.loadFavoriteds(ctx, query, nodes,
			func(n *Account) { n.Edges.Favoriteds = []*Account{} },
			func(n *Account, e *Account) { n.Edges.Favoriteds = append(n.Edges.Favoriteds, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withFavorites; query != nil {
		if err := aq.loadFavorites(ctx, query, nodes,
			func(n *Account) { n.Edges.Favorites = []*Account{} },
			func(n *Account, e *Account) { n.Edges.Favorites = append(n.Edges.Favorites, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withFromAccount; query != nil {
		if err := aq.loadFromAccount(ctx, query, nodes,
			func(n *Account) { n.Edges.FromAccount = []*Transaction{} },
			func(n *Account, e *Transaction) { n.Edges.FromAccount = append(n.Edges.FromAccount, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withToAccount; query != nil {
		if err := aq.loadToAccount(ctx, query, nodes,
			func(n *Account) { n.Edges.ToAccount = []*Transaction{} },
			func(n *Account, e *Transaction) { n.Edges.ToAccount = append(n.Edges.ToAccount, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withDepositRequest; query != nil {
		if err := aq.loadDepositRequest(ctx, query, nodes,
			func(n *Account) { n.Edges.DepositRequest = []*DepositRequest{} },
			func(n *Account, e *DepositRequest) { n.Edges.DepositRequest = append(n.Edges.DepositRequest, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withNotification; query != nil {
		if err := aq.loadNotification(ctx, query, nodes,
			func(n *Account) { n.Edges.Notification = []*Notification{} },
			func(n *Account, e *Notification) { n.Edges.Notification = append(n.Edges.Notification, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AccountQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Account, init func(*Account), assign func(*Account, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(account.UserColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.account_user
		if fk == nil {
			return fmt.Errorf(`foreign-key "account_user" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "account_user" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AccountQuery) loadFavoriteds(ctx context.Context, query *AccountQuery, nodes []*Account, init func(*Account), assign func(*Account, *Account)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Account)
	nids := make(map[string]map[*Account]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(account.FavoritedsTable)
		s.Join(joinT).On(s.C(account.FieldID), joinT.C(account.FavoritedsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(account.FavoritedsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(account.FavoritedsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Account]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Account](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "favoriteds" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AccountQuery) loadFavorites(ctx context.Context, query *AccountQuery, nodes []*Account, init func(*Account), assign func(*Account, *Account)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Account)
	nids := make(map[string]map[*Account]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(account.FavoritesTable)
		s.Join(joinT).On(s.C(account.FieldID), joinT.C(account.FavoritesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(account.FavoritesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(account.FavoritesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Account]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Account](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "favorites" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AccountQuery) loadFromAccount(ctx context.Context, query *TransactionQuery, nodes []*Account, init func(*Account), assign func(*Account, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(account.FromAccountColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.account_from_account
		if fk == nil {
			return fmt.Errorf(`foreign-key "account_from_account" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "account_from_account" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AccountQuery) loadToAccount(ctx context.Context, query *TransactionQuery, nodes []*Account, init func(*Account), assign func(*Account, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(account.ToAccountColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.account_to_account
		if fk == nil {
			return fmt.Errorf(`foreign-key "account_to_account" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "account_to_account" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AccountQuery) loadDepositRequest(ctx context.Context, query *DepositRequestQuery, nodes []*Account, init func(*Account), assign func(*Account, *DepositRequest)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DepositRequest(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(account.DepositRequestColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.account_deposit_request
		if fk == nil {
			return fmt.Errorf(`foreign-key "account_deposit_request" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "account_deposit_request" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AccountQuery) loadNotification(ctx context.Context, query *NotificationQuery, nodes []*Account, init func(*Account), assign func(*Account, *Notification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(account.NotificationColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.account_notification
		if fk == nil {
			return fmt.Errorf(`foreign-key "account_notification" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "account_notification" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeString))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for i := range fields {
			if fields[i] != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(account.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = account.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountGroupBy is the group-by builder for Account entities.
type AccountGroupBy struct {
	selector
	build *AccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AccountGroupBy) Aggregate(fns ...AggregateFunc) *AccountGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountQuery, *AccountGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AccountGroupBy) sqlScan(ctx context.Context, root *AccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountSelect is the builder for selecting fields of Account entities.
type AccountSelect struct {
	*AccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AccountSelect) Aggregate(fns ...AggregateFunc) *AccountSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountQuery, *AccountSelect](ctx, as.AccountQuery, as, as.inters, v)
}

func (as *AccountSelect) sqlScan(ctx context.Context, root *AccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
