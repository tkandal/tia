package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// Name represents a name in TIA
type Name struct {
	TiaPersistable
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

// MakeName returns a new Name struct.
func MakeName(typ string, value string, lang string) Name {
	n := Name{
		Type:  typ,
		Value: value,
	}
	if len(lang) > 0 {
		n.Language = lang
	}
	return n
}

// NewName allocates a new Name-object and returns a pointer to the new Name-object.
func NewName(typ string, value string, lang string) *Name {
	n := MakeName(typ, value, lang)
	return &n
}
