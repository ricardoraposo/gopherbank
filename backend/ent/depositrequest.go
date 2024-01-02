// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ricardoraposo/gopherbank/ent/account"
	"github.com/ricardoraposo/gopherbank/ent/depositrequest"
)

// DepositRequest is the model entity for the DepositRequest schema.
type DepositRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount"`
	// Status holds the value of the "status" field.
	Status string `json:"status"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepositRequestQuery when eager-loading is set.
	Edges                   DepositRequestEdges `json:"edges"`
	account_deposit_request *string
	selectValues            sql.SelectValues
}

// DepositRequestEdges holds the relations/edges for other nodes in the graph.
type DepositRequestEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepositRequestEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DepositRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case depositrequest.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case depositrequest.FieldID:
			values[i] = new(sql.NullInt64)
		case depositrequest.FieldStatus:
			values[i] = new(sql.NullString)
		case depositrequest.ForeignKeys[0]: // account_deposit_request
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DepositRequest fields.
func (dr *DepositRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case depositrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dr.ID = int(value.Int64)
		case depositrequest.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				dr.Amount = value.Float64
			}
		case depositrequest.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dr.Status = value.String
			}
		case depositrequest.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_deposit_request", values[i])
			} else if value.Valid {
				dr.account_deposit_request = new(string)
				*dr.account_deposit_request = value.String
			}
		default:
			dr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DepositRequest.
// This includes values selected through modifiers, order, etc.
func (dr *DepositRequest) Value(name string) (ent.Value, error) {
	return dr.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the DepositRequest entity.
func (dr *DepositRequest) QueryAccount() *AccountQuery {
	return NewDepositRequestClient(dr.config).QueryAccount(dr)
}

// Update returns a builder for updating this DepositRequest.
// Note that you need to call DepositRequest.Unwrap() before calling this method if this DepositRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DepositRequest) Update() *DepositRequestUpdateOne {
	return NewDepositRequestClient(dr.config).UpdateOne(dr)
}

// Unwrap unwraps the DepositRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dr *DepositRequest) Unwrap() *DepositRequest {
	_tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DepositRequest is not a transactional entity")
	}
	dr.config.driver = _tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DepositRequest) String() string {
	var builder strings.Builder
	builder.WriteString("DepositRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dr.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", dr.Amount))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(dr.Status)
	builder.WriteByte(')')
	return builder.String()
}

// DepositRequests is a parsable slice of DepositRequest.
type DepositRequests []*DepositRequest