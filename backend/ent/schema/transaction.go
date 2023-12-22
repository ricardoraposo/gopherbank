package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
    return nil
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("from_account", Account.Type).
            Unique().
            StorageKey(edge.Column("from_account_number")).
            StructTag(`json:"fromAccount"`),
        edge.To("to_account", Account.Type).
            Unique().
            StorageKey(edge.Column("to_account_number")).
            StructTag(`json:"toAccount"`),
        edge.To("detail", TransactionDetail.Type).
            StorageKey(edge.Column("transaction_id")).
            Unique(),
    }
}
