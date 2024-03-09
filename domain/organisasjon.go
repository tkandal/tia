package domain

import "time"

/*
 * Copyright (c) 2021, 2022 Norwegian University of Science and Technology
 */

// Organisasjon represents an organisation in TIA.
// swagger:model Organisasjon
type Organisasjon struct {
	Entity
	Name        []Name        `json:"name,omitempty" bson:"name,omitempty"`
	ContactInfo []ContactInfo `json:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	Status      []Status      `json:"status,omitempty" bson:"status,omitempty"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	ValidFrom   *time.Time    `json:"validFrom,omitempty" bson:"validFrom,omitempty"`
	ValidTo     *time.Time    `json:"validTo,omitempty" bson:"validTo,omitempty"`
}
