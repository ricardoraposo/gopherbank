package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        // field.String("id").StorageKey("user_account").StructTag(`json:"account"`),
        field.String("first_name").MaxLen(50).StructTag(`json:"firstName"`),
        field.String("last_name").MaxLen(50).StructTag(`json:"lastName"`),
        field.String("email").StructTag(`json:"email"`),
        field.String("picture_url").StructTag(`json:"pictureUrl"`),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("account", Account.Type).Ref("user").Unique().Required(),
    }
}
