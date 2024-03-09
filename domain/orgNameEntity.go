package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
// Deprecated: These objects has been replaced, and should not be used.
type OrgNameEntity struct {
	TiaPersistable
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
	Acronym   string     `json:"acronym,omitempty" bson:"acronym,omitempty"`
	Language  string     `json:"language,omitempty" bson:"language,omitempty"`
	ValidFrom *time.Time `json:"validFrom,omitempty" bson:"validFrom,omitempty"`
	Expires   *time.Time `json:"expires,omitempty" bson:"expires,omitempty"`
}
