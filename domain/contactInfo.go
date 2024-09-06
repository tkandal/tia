package domain

import "sort"

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
func MakeContactInfo(typ string, value string, lang string) ContactInfo {
	ci := ContactInfo{
		Type:  typ,
		Value: value,
	}
	if len(lang) > 0 {
		ci.Language = lang
	}
	return ci
}

// NewContactInfo returns a pointer to a new contact-information struct
func NewContactInfo(typ string, value string, lang string) *ContactInfo {
	ci := MakeContactInfo(typ, value, lang)
	return &ci
}

type ContactInfoBuilder struct {
	ci *ContactInfo
}

func NewContactInfoBuilder(typ string) *ContactInfoBuilder {
	return &ContactInfoBuilder{
		ci: &ContactInfo{
			Type: typ,
		},
	}
}

func (b *ContactInfoBuilder) Add(ext string) *ContactInfoBuilder {
	if len(ext) > 0 {
		if b.ci.AddressInfo == nil {
			b.ci.AddressInfo = make([]string, 0)
		}
		if !contains(b.ci.AddressInfo, ext) {
			b.ci.AddressInfo = append(b.ci.AddressInfo, ext)
		}
	}
	return b
}

func (b *ContactInfoBuilder) PostalCode(code string) *ContactInfoBuilder {
	if len(code) > 0 {
		b.ci.PostalCode = code
	}
	return b
}

func (b *ContactInfoBuilder) City(city string) *ContactInfoBuilder {
	if len(city) > 0 {
		b.ci.City = city
	}
	return b
}

func (b *ContactInfoBuilder) State(state string) *ContactInfoBuilder {
	if len(state) > 0 {
		b.ci.State = state
	}
	return b
}

func (b *ContactInfoBuilder) Country(country string) *ContactInfoBuilder {
	if len(country) > 0 {
		b.ci.Country = country
	}
	return b
}

func (b *ContactInfoBuilder) Priority(p int64) *ContactInfoBuilder {
	b.ci.Priority = p
	return b
}

func (b *ContactInfoBuilder) Room(room string) *ContactInfoBuilder {
	if len(room) > 0 {
		b.ci.Room = room
	}
	return b
}

func (b *ContactInfoBuilder) Building(building string) *ContactInfoBuilder {
	if len(building) > 0 {
		b.ci.Building = building
	}
	return b
}

func (b *ContactInfoBuilder) LydiaCode(lc string) *ContactInfoBuilder {
	if len(lc) == 0 {
		b.ci.LydiaCode = lc
	}
	return b
}

func (b *ContactInfoBuilder) CountryNo(cn *EntityKey) *ContactInfoBuilder {
	if cn != nil {
		b.ci.CountryNo = cn
	}
	return b
}

func (b *ContactInfoBuilder) CountyNo(cn *EntityKey) *ContactInfoBuilder {
	if cn != nil {
		b.ci.CountyNo = cn
	}
	return b
}

func (b *ContactInfoBuilder) MunicipalNo(mn *EntityKey) *ContactInfoBuilder {
	if mn != nil {
		b.ci.MunicipalNo = mn
	}
	return b
}

func (b *ContactInfoBuilder) PhoneCountryNo(pn string) *ContactInfoBuilder {
	if len(pn) > 0 {
		b.ci.PhoneCountryNo = pn
	}
	return b
}

func (b *ContactInfoBuilder) Build() *ContactInfo {
	return b.ci
}

func contains(stack []string, needle string) bool {
	if len(stack) == 0 {
		return false
	}
	sort.Strings(stack)
	idx := sort.SearchStrings(stack, needle)
	return idx < len(stack) && stack[idx] == needle
}
