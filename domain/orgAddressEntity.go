package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */

// Deprecated: These objects has been replaced, and should not be used.
type OrgAddressEntity struct {
	TiaPersistable
	AddressClass string `json:"addressClass,omitempty" bson:"addressClass,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty" bson:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty" bson:"addressLine2,omitempty"`
	City         string `json:"city,omitempty" bson:"city,omitempty"`
	Country      string `json:"country,omitempty" bson:"country,omitempty"`
	PostalCode   string `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	Language     string `json:"language,omitempty" bson:"language,omitempty"`
}
