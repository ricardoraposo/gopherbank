package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TransactionDetail holds the schema definition for the TransactionDetail entity.
type TransactionDetail struct {
	ent.Schema
}

// Fields of the TransactionDetail.
func (TransactionDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("transaction_id").StructTag(`json:"transactionId"`),
		field.Float("amount").StructTag(`json:"amount"`),
		field.String("type").MaxLen(20).StructTag(`json:"type"`),
		field.Time("transacted_at").Default(time.Now).Immutable().StructTag(`json:"transactedAt"`),
	}
}

// Edges of the TransactionDetail.
func (TransactionDetail) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("transaction", Transaction.Type).
            Ref("detail").
            Unique().
            StructTag(`json:"transactionId"`),
    }
}
