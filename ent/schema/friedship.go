package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Friendship struct {
	ent.Schema
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
		field.Int("user_id"),
		field.Int("friend_id"),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("friend", User.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}

// Annotations of the Friendship.
func (Friendship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
