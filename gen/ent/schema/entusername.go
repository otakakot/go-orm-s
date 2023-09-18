package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntUserName holds the schema definition for the EntUserName entity.
type EntUserName struct {
	ent.Schema
}

// Fields of the EntUserName.
func (EntUserName) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("value"),
		field.Time("created_at").Default(time.Now().UTC),
		field.Time("updated_at").Default(time.Now().UTC).UpdateDefault(time.Now().UTC),
		field.Bool("deleted").Default(false),
	}
}

// Edges of the EntUserName.
func (EntUserName) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ent_user", EntUser.Type).Ref("ent_user_names").Field("user_id").Required().Unique(),
	}
}
