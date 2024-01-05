package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("number").Unique().StructTag(`json:"number"`),
		field.String("password").StructTag(`json:"-"`),
		field.Float("balance").Default(0).StructTag(`json:"balance"`),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Bool("admin").Default(false).StructTag(`json:"admin"`),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("user", User.Type).Unique(),
        edge.To("favorites", Account.Type).From("favoriteds"),
        edge.To("from_account", Transaction.Type).StructTag(`json:"fromAccount"`),
        edge.To("to_account", Transaction.Type).StructTag(`json:"toAccount"`),
        edge.To("deposit_request", DepositRequest.Type),
        edge.To("notification", Notification.Type),
    }
}
