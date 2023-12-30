// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ricardoraposo/gopherbank/ent/transaction"
	"github.com/ricardoraposo/gopherbank/ent/transactiondetails"
)

// TransactionDetails is the model entity for the TransactionDetails schema.
type TransactionDetails struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount"`
	// Type holds the value of the "type" field.
	Type string `json:"type"`
	// TransactedAt holds the value of the "transacted_at" field.
	TransactedAt time.Time `json:"transactedAt"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionDetailsQuery when eager-loading is set.
	Edges              TransactionDetailsEdges `json:"edges"`
	transaction_detail *int
	selectValues       sql.SelectValues
}

// TransactionDetailsEdges holds the relations/edges for other nodes in the graph.
type TransactionDetailsEdges struct {
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionDetailsEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[0] {
		if e.Transaction == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransactionDetails) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transactiondetails.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case transactiondetails.FieldID:
			values[i] = new(sql.NullInt64)
		case transactiondetails.FieldType:
			values[i] = new(sql.NullString)
		case transactiondetails.FieldTransactedAt:
			values[i] = new(sql.NullTime)
		case transactiondetails.ForeignKeys[0]: // transaction_detail
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransactionDetails fields.
func (td *TransactionDetails) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transactiondetails.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			td.ID = int(value.Int64)
		case transactiondetails.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				td.Amount = value.Float64
			}
		case transactiondetails.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				td.Type = value.String
			}
		case transactiondetails.FieldTransactedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field transacted_at", values[i])
			} else if value.Valid {
				td.TransactedAt = value.Time
			}
		case transactiondetails.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_detail", value)
			} else if value.Valid {
				td.transaction_detail = new(int)
				*td.transaction_detail = int(value.Int64)
			}
		default:
			td.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TransactionDetails.
// This includes values selected through modifiers, order, etc.
func (td *TransactionDetails) Value(name string) (ent.Value, error) {
	return td.selectValues.Get(name)
}

// QueryTransaction queries the "transaction" edge of the TransactionDetails entity.
func (td *TransactionDetails) QueryTransaction() *TransactionQuery {
	return NewTransactionDetailsClient(td.config).QueryTransaction(td)
}

// Update returns a builder for updating this TransactionDetails.
// Note that you need to call TransactionDetails.Unwrap() before calling this method if this TransactionDetails
// was returned from a transaction, and the transaction was committed or rolled back.
func (td *TransactionDetails) Update() *TransactionDetailsUpdateOne {
	return NewTransactionDetailsClient(td.config).UpdateOne(td)
}

// Unwrap unwraps the TransactionDetails entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (td *TransactionDetails) Unwrap() *TransactionDetails {
	_tx, ok := td.config.driver.(*txDriver)
	if !ok {
		panic("ent: TransactionDetails is not a transactional entity")
	}
	td.config.driver = _tx.drv
	return td
}

// String implements the fmt.Stringer.
func (td *TransactionDetails) String() string {
	var builder strings.Builder
	builder.WriteString("TransactionDetails(")
	builder.WriteString(fmt.Sprintf("id=%v, ", td.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", td.Amount))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(td.Type)
	builder.WriteString(", ")
	builder.WriteString("transacted_at=")
	builder.WriteString(td.TransactedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TransactionDetailsSlice is a parsable slice of TransactionDetails.
type TransactionDetailsSlice []*TransactionDetails
