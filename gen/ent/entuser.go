// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/otakakot/go-orm-s/gen/ent/entuser"
)

// EntUser is the model entity for the EntUser schema.
type EntUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntUserQuery when eager-loading is set.
	Edges        EntUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EntUserEdges holds the relations/edges for other nodes in the graph.
type EntUserEdges struct {
	// EntUserNames holds the value of the ent_user_names edge.
	EntUserNames []*EntUserName `json:"ent_user_names,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EntUserNamesOrErr returns the EntUserNames value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) EntUserNamesOrErr() ([]*EntUserName, error) {
	if e.loadedTypes[0] {
		return e.EntUserNames, nil
	}
	return nil, &NotLoadedError{edge: "ent_user_names"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entuser.FieldDeleted:
			values[i] = new(sql.NullBool)
		case entuser.FieldCreatedAt, entuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case entuser.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntUser fields.
func (eu *EntUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				eu.ID = *value
			}
		case entuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eu.CreatedAt = value.Time
			}
		case entuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eu.UpdatedAt = value.Time
			}
		case entuser.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				eu.Deleted = value.Bool
			}
		default:
			eu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EntUser.
// This includes values selected through modifiers, order, etc.
func (eu *EntUser) Value(name string) (ent.Value, error) {
	return eu.selectValues.Get(name)
}

// QueryEntUserNames queries the "ent_user_names" edge of the EntUser entity.
func (eu *EntUser) QueryEntUserNames() *EntUserNameQuery {
	return NewEntUserClient(eu.config).QueryEntUserNames(eu)
}

// Update returns a builder for updating this EntUser.
// Note that you need to call EntUser.Unwrap() before calling this method if this EntUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (eu *EntUser) Update() *EntUserUpdateOne {
	return NewEntUserClient(eu.config).UpdateOne(eu)
}

// Unwrap unwraps the EntUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eu *EntUser) Unwrap() *EntUser {
	_tx, ok := eu.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntUser is not a transactional entity")
	}
	eu.config.driver = _tx.drv
	return eu
}

// String implements the fmt.Stringer.
func (eu *EntUser) String() string {
	var builder strings.Builder
	builder.WriteString("EntUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eu.ID))
	builder.WriteString("created_at=")
	builder.WriteString(eu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(eu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", eu.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// EntUsers is a parsable slice of EntUser.
type EntUsers []*EntUser
