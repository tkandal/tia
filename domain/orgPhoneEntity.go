package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
// Deprecated: These objects has been replaced, and should not be used.
type OrgPhoneEntity struct {
	TiaPersistable
	PhoneType string `json:"phoneType,omitempty" bson:"phoneType,omitempty"`
	Number    string `json:"number,omitempty" bson:"number,omitempty"`
	Prefix    string `json:"prefix,omitempty" bson:"prefix,omitempty"`
}
