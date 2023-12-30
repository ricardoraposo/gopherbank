// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "number"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeFavoriteds holds the string denoting the favoriteds edge name in mutations.
	EdgeFavoriteds = "favoriteds"
	// EdgeFavorites holds the string denoting the favorites edge name in mutations.
	EdgeFavorites = "favorites"
	// EdgeFromAccount holds the string denoting the from_account edge name in mutations.
	EdgeFromAccount = "from_account"
	// EdgeToAccount holds the string denoting the to_account edge name in mutations.
	EdgeToAccount = "to_account"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// TransactionFieldID holds the string denoting the ID field of the Transaction.
	TransactionFieldID = "id"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "account_user"
	// FavoritedsTable is the table that holds the favoriteds relation/edge. The primary key declared below.
	FavoritedsTable = "account_favorites"
	// FavoritesTable is the table that holds the favorites relation/edge. The primary key declared below.
	FavoritesTable = "account_favorites"
	// FromAccountTable is the table that holds the from_account relation/edge.
	FromAccountTable = "transactions"
	// FromAccountInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	FromAccountInverseTable = "transactions"
	// FromAccountColumn is the table column denoting the from_account relation/edge.
	FromAccountColumn = "account_from_account"
	// ToAccountTable is the table that holds the to_account relation/edge.
	ToAccountTable = "transactions"
	// ToAccountInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	ToAccountInverseTable = "transactions"
	// ToAccountColumn is the table column denoting the to_account relation/edge.
	ToAccountColumn = "account_to_account"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldPassword,
	FieldBalance,
	FieldCreatedAt,
	FieldAdmin,
}

var (
	// FavoritedsPrimaryKey and FavoritedsColumn2 are the table columns denoting the
	// primary key for the favoriteds relation (M2M).
	FavoritedsPrimaryKey = []string{"account_id", "favorited_id"}
	// FavoritesPrimaryKey and FavoritesColumn2 are the table columns denoting the
	// primary key for the favorites relation (M2M).
	FavoritesPrimaryKey = []string{"account_id", "favorited_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance float64
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
)

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAdmin orders the results by the admin field.
func ByAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdmin, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByFavoritedsCount orders the results by favoriteds count.
func ByFavoritedsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFavoritedsStep(), opts...)
	}
}

// ByFavoriteds orders the results by favoriteds terms.
func ByFavoriteds(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavoritedsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFavoritesCount orders the results by favorites count.
func ByFavoritesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFavoritesStep(), opts...)
	}
}

// ByFavorites orders the results by favorites terms.
func ByFavorites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavoritesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFromAccountCount orders the results by from_account count.
func ByFromAccountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFromAccountStep(), opts...)
	}
}

// ByFromAccount orders the results by from_account terms.
func ByFromAccount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromAccountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByToAccountCount orders the results by to_account count.
func ByToAccountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newToAccountStep(), opts...)
	}
}

// ByToAccount orders the results by to_account terms.
func ByToAccount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToAccountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, UserTable, UserColumn),
	)
}
func newFavoritedsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, FavoritedsTable, FavoritedsPrimaryKey...),
	)
}
func newFavoritesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FavoritesTable, FavoritesPrimaryKey...),
	)
}
func newFromAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromAccountInverseTable, TransactionFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FromAccountTable, FromAccountColumn),
	)
}
func newToAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToAccountInverseTable, TransactionFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ToAccountTable, ToAccountColumn),
	)
}
