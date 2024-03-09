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

// MakeName returns a new Name struct
func MakeName(typ string, value string) Name {
	return Name{
		Type:  typ,
		Value: value,
	}
}
