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
        edge.From("from_account", Account.Type).Ref("from_account").Unique(),
        edge.From("to_account", Account.Type).Ref("to_account").Unique(),
        edge.To("detail", TransactionDetails.Type).Unique(),
    }
}
