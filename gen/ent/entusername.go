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
	"github.com/otakakot/go-orm-s/gen/ent/entusername"
)

// EntUserName is the model entity for the EntUserName schema.
type EntUserName struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntUserNameQuery when eager-loading is set.
	Edges        EntUserNameEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EntUserNameEdges holds the relations/edges for other nodes in the graph.
type EntUserNameEdges struct {
	// EntUser holds the value of the ent_user edge.
	EntUser *EntUser `json:"ent_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EntUserOrErr returns the EntUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntUserNameEdges) EntUserOrErr() (*EntUser, error) {
	if e.loadedTypes[0] {
		if e.EntUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: entuser.Label}
		}
		return e.EntUser, nil
	}
	return nil, &NotLoadedError{edge: "ent_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntUserName) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entusername.FieldDeleted:
			values[i] = new(sql.NullBool)
		case entusername.FieldValue:
			values[i] = new(sql.NullString)
		case entusername.FieldCreatedAt, entusername.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case entusername.FieldID, entusername.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntUserName fields.
func (eun *EntUserName) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entusername.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				eun.ID = *value
			}
		case entusername.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				eun.UserID = *value
			}
		case entusername.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				eun.Value = value.String
			}
		case entusername.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eun.CreatedAt = value.Time
			}
		case entusername.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eun.UpdatedAt = value.Time
			}
		case entusername.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				eun.Deleted = value.Bool
			}
		default:
			eun.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the EntUserName.
// This includes values selected through modifiers, order, etc.
func (eun *EntUserName) GetValue(name string) (ent.Value, error) {
	return eun.selectValues.Get(name)
}

// QueryEntUser queries the "ent_user" edge of the EntUserName entity.
func (eun *EntUserName) QueryEntUser() *EntUserQuery {
	return NewEntUserNameClient(eun.config).QueryEntUser(eun)
}

// Update returns a builder for updating this EntUserName.
// Note that you need to call EntUserName.Unwrap() before calling this method if this EntUserName
// was returned from a transaction, and the transaction was committed or rolled back.
func (eun *EntUserName) Update() *EntUserNameUpdateOne {
	return NewEntUserNameClient(eun.config).UpdateOne(eun)
}

// Unwrap unwraps the EntUserName entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eun *EntUserName) Unwrap() *EntUserName {
	_tx, ok := eun.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntUserName is not a transactional entity")
	}
	eun.config.driver = _tx.drv
	return eun
}

// String implements the fmt.Stringer.
func (eun *EntUserName) String() string {
	var builder strings.Builder
	builder.WriteString("EntUserName(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eun.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", eun.UserID))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(eun.Value)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(eun.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(eun.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", eun.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// EntUserNames is a parsable slice of EntUserName.
type EntUserNames []*EntUserName
