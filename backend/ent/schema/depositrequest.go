package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DepositRequest holds the schema definition for the DepositRequest entity.
type DepositRequest struct {
	ent.Schema
}

// Fields of the DepositRequest.
func (DepositRequest) Fields() []ent.Field {
    status := regexp.MustCompile(`^(pending|approved|rejected)$`)

	return []ent.Field {
        field.Float("amount").Positive().StructTag(`json:"amount"`),
        field.String("status").Default("pending").Match(status).StructTag(`json:"status"`),
    }
}

// Edges of the DepositRequest.
func (DepositRequest) Edges() []ent.Edge {
	return []ent.Edge {
        edge.From("account", Account.Type).Ref("deposit_request").Unique(),
    }
}
