package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
        field.String("title").NotEmpty(),
        field.String("content").NotEmpty(),
        field.Bool("read").Default(false),
        field.Time("created_at").Default(time.Now).Immutable().StructTag(`json:"createdAt"`),
    }
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("account", Account.Type).Ref("notification").Unique(),
    }
}
