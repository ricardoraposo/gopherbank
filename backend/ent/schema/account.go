package schema

import (
	"time"

	"entgo.io/ent"
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
		field.String("password").StructTag(`json:"password"`),
		field.Float("balance").StructTag(`json:"balance"`),
		field.Time("createdAt").Default(time.Now).Immutable().StructTag(`json:"createdAt"`),
		field.Bool("admin").Default(false).StructTag(`json:"admin"`),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
