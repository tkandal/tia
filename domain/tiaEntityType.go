package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// TiaEntityType is the type of the entity
type TiaEntityType struct {
	TiaPersistable
	EntityTypeID string `json:"entityTypeId,omitempty" bson:"entityTypeId,omitempty"`
}

// MakeTiaEntityType returns a new TiaEntityType struct
func MakeTiaEntityType(ei string) TiaEntityType {
	return TiaEntityType{
		EntityTypeID: ei,
	}
}

// NewTiaEntityType returns a pointer a new TiaEntityType struct
func NewTiaEntityType(ei string) *TiaEntityType {
	return &TiaEntityType{EntityTypeID: ei}
}
