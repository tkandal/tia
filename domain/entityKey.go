package domain

import (
	"fmt"
)

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// EntityKey represents a key-value and -type in TIA
type EntityKey struct {
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	Value       string `json:"value,omitempty" bson:"value,omitempty"`
	EntityKeyID string `json:"entityKeyId,omitempty" bson:"entityKeyId,omitempty"`
}

// MakeEntityKey returns a new EntityKey struct
func MakeEntityKey(typ string, value string) EntityKey {
	id := fmt.Sprintf("%s:%s", typ, value)
	return EntityKey{
		Type:        typ,
		Value:       value,
		EntityKeyID: id,
	}
}

// NewEntityKey returns a pointer to a new EntityKey struct
func NewEntityKey(typ string, value string) *EntityKey {
	id := fmt.Sprintf("%s:%s", typ, value)
	return &EntityKey{
		Type:        typ,
		Value:       value,
		EntityKeyID: id,
	}
}
