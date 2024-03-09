package domain

import (
	"time"
)

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// TiaSuper is basic structures which very structure inherits
type TiaSuper struct {
	LastUpdated *time.Time `json:"lastUpdated,omitempty" bson:"lastUpdated,omitempty"`
	ValidFrom   *time.Time `json:"validFrom,omitempty" bson:"validFrom,omitempty"`
	ValidTo     *time.Time `json:"validTo,omitempty" bson:"validTo,omitempty"`
	Language    string     `json:"language,omitempty" bson:"language,omitempty"`
}
