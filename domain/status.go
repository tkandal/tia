package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// Status holds the type and value for a status
type Status struct {
	TiaPersistable
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

// MakeStatus return a new Status struct
func MakeStatus(typ string, value string) Status {
	return Status{
		Type:  typ,
		Value: value,
	}
}
