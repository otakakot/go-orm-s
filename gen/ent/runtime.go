// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/otakakot/go-orm-s/gen/ent/entuser"
	"github.com/otakakot/go-orm-s/gen/ent/entusername"
	"github.com/otakakot/go-orm-s/gen/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entuserFields := schema.EntUser{}.Fields()
	_ = entuserFields
	// entuserDescCreatedAt is the schema descriptor for created_at field.
	entuserDescCreatedAt := entuserFields[1].Descriptor()
	// entuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	entuser.DefaultCreatedAt = entuserDescCreatedAt.Default.(func() time.Time)
	// entuserDescUpdatedAt is the schema descriptor for updated_at field.
	entuserDescUpdatedAt := entuserFields[2].Descriptor()
	// entuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entuser.DefaultUpdatedAt = entuserDescUpdatedAt.Default.(func() time.Time)
	// entuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entuser.UpdateDefaultUpdatedAt = entuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// entuserDescDeleted is the schema descriptor for deleted field.
	entuserDescDeleted := entuserFields[3].Descriptor()
	// entuser.DefaultDeleted holds the default value on creation for the deleted field.
	entuser.DefaultDeleted = entuserDescDeleted.Default.(bool)
	// entuserDescID is the schema descriptor for id field.
	entuserDescID := entuserFields[0].Descriptor()
	// entuser.DefaultID holds the default value on creation for the id field.
	entuser.DefaultID = entuserDescID.Default.(func() uuid.UUID)
	entusernameFields := schema.EntUserName{}.Fields()
	_ = entusernameFields
	// entusernameDescCreatedAt is the schema descriptor for created_at field.
	entusernameDescCreatedAt := entusernameFields[3].Descriptor()
	// entusername.DefaultCreatedAt holds the default value on creation for the created_at field.
	entusername.DefaultCreatedAt = entusernameDescCreatedAt.Default.(func() time.Time)
	// entusernameDescUpdatedAt is the schema descriptor for updated_at field.
	entusernameDescUpdatedAt := entusernameFields[4].Descriptor()
	// entusername.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entusername.DefaultUpdatedAt = entusernameDescUpdatedAt.Default.(func() time.Time)
	// entusername.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entusername.UpdateDefaultUpdatedAt = entusernameDescUpdatedAt.UpdateDefault.(func() time.Time)
	// entusernameDescDeleted is the schema descriptor for deleted field.
	entusernameDescDeleted := entusernameFields[5].Descriptor()
	// entusername.DefaultDeleted holds the default value on creation for the deleted field.
	entusername.DefaultDeleted = entusernameDescDeleted.Default.(bool)
	// entusernameDescID is the schema descriptor for id field.
	entusernameDescID := entusernameFields[0].Descriptor()
	// entusername.DefaultID holds the default value on creation for the id field.
	entusername.DefaultID = entusernameDescID.Default.(func() uuid.UUID)
}