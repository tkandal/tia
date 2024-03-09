package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// ContactInfo is a person's contact-information
type ContactInfo struct {
	TiaPersistable
	Type           string     `json:"type,omitempty" bson:"type,omitempty"`
	Value          string     `json:"value,omitempty" bson:"value,omitempty"`
	AddressInfo    []string   `json:"addressInfo,omitempty" bson:"addressInfo,omitempty"`
	PostalCode     string     `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	City           string     `json:"city,omitempty" bson:"city,omitempty"`
	State          string     `json:"state,omitempty" bson:"state,omitempty"`
	Country        string     `json:"country,omitempty" bson:"country,omitempty"`
	Priority       int64      `json:"priority,omitempty" bson:"priority,omitempty"`
	Room           string     `json:"room,omitempty" bson:"room,omitempty"`
	Building       string     `json:"building,omitempty" bson:"building,omitempty"`
	LydiaCode      string     `json:"lydiacode,omitempty" bson:"lydiacode,omitempty"`
	CountryNo      *EntityKey `json:"countryNo,omitempty" bson:"countryNo,omitempty"`
	CountyNo       *EntityKey `json:"countyNo,omitempty" bson:"countyNo,omitempty"`
	MunicipalNo    *EntityKey `json:"municipalNo,omitempty" bson:"municipalNo,omitempty"`
	PhoneCountryNo string     `json:"phoneCountryNo,omitempty" bson:"phoneCountryNo,omitempty"`
}

// MakeContactInfo makes a new contact-information struct
func MakeContactInfo(typ string, value string) ContactInfo {
	return ContactInfo{
		Type:  typ,
		Value: value,
	}
}

// NewContactInfo returns a pointer to a new contact-information struct
func NewContactInfo(typ string, value string) *ContactInfo {
	return &ContactInfo{
		Type:  typ,
		Value: value,
	}
}
