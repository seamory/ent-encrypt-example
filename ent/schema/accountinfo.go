package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"main/internal/property"
)

// AccountInfo holds the schema definition for the AccountInfo entity.
type AccountInfo struct {
	ent.Schema
}

// Fields of the AccountInfo.
func (AccountInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password").GoType(property.Password("")),
	}
}

// Edges of the AccountInfo.
func (AccountInfo) Edges() []ent.Edge {
	return nil
}
