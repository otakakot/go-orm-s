package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntUser holds the schema definition for the EntUser entity.
type EntUser struct {
	ent.Schema
}

// Fields of the EntUser.
func (EntUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now().UTC),
		field.Time("updated_at").Default(time.Now().UTC).UpdateDefault(time.Now().UTC),
		field.Bool("deleted").Default(false),
	}
}

// Edges of the EntUser.
func (EntUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ent_user_names", EntUserName.Type),
	}
}
