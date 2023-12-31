// Code generated by ent, DO NOT EDIT.

package transactiondetails

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the transactiondetails type in the database.
	Label = "transaction_details"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTransactedAt holds the string denoting the transacted_at field in the database.
	FieldTransactedAt = "transacted_at"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"
	// Table holds the table name of the transactiondetails in the database.
	Table = "transaction_details"
	// TransactionTable is the table that holds the transaction relation/edge.
	TransactionTable = "transaction_details"
	// TransactionInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionInverseTable = "transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "transaction_detail"
)

// Columns holds all SQL columns for transactiondetails fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldType,
	FieldTransactedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transaction_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"transaction_detail",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultTransactedAt holds the default value on creation for the "transacted_at" field.
	DefaultTransactedAt func() time.Time
)

// OrderOption defines the ordering options for the TransactionDetails queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByTransactedAt orders the results by the transacted_at field.
func ByTransactedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactedAt, opts...).ToFunc()
}

// ByTransactionField orders the results by transaction field.
func ByTransactionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionStep(), sql.OrderByField(field, opts...))
	}
}
func newTransactionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, TransactionTable, TransactionColumn),
	)
}
