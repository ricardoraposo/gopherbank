package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TransactionDetails holds the schema definition for the TransactionDetails entity.
type TransactionDetails struct {
	ent.Schema
}

// Fields of the TransactionDetails.
func (TransactionDetails) Fields() []ent.Field {
	types := regexp.MustCompile("^(deposit|withdraw|transfer)$")

	return []ent.Field{
		field.Int("id").StorageKey("transaction_id").StructTag(`json:"transactionId"`),
		field.Float("amount").StructTag(`json:"amount"`),
		field.String("type").Match(types).StructTag(`json:"type"`),
		field.Time("transacted_at").Default(time.Now).Immutable().StructTag(`json:"transactedAt"`),
	}
}

// Edges of the TransactionDetails.
func (TransactionDetails) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("transaction", Transaction.Type).
            Ref("detail").
            Unique().
            Required(),
    }
}
