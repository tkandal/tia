package domain

import (
	"encoding/json"
	"strings"
	"time"
)

/*
 * Copyright (c) 2019, 2022 Norwegian University of Science and Technology
 */

// Person represents a person in TIA.
// swagger:model Person
type Person struct {
	Entity
	Name              []Name        `json:"name,omitempty" bson:"name,omitempty"`
	ContactInfo       []ContactInfo `json:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	Cars              []Car         `json:"cars,omitempty" bson:"cars,omitempty"`
	Status            []Status      `json:"status,omitempty" bson:"status,omitempty"`
	DateOfBirth       *LocalDate    `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	Gender            string        `json:"gender,omitempty" bson:"gender,omitempty"`
	PreferredLanguage string        `json:"preferredLanguage,omitempty" bson:"preferredLanguage,omitempty"`
	NativeTongue      []string      `json:"nativeTongue,omitempty" bson:"nativeTongue,omitempty"`
	SpecialNeed       []string      `json:"specialNeed,omitempty" bson:"specialNeed,omitempty"`
}

func (p *Person) String() string {
	str := &strings.Builder{}
	if err := json.NewEncoder(str).Encode(p); err != nil {
		return ""
	}
	return str.String()
}

// PersonBuilder is a struct for person-builder
type PersonBuilder struct {
	pers *Person
}

// NewPersonBuilder returns a pointer to a person-builder
func NewPersonBuilder(ss string, ssID string) (*PersonBuilder, error) {
	pers := &Person{}
	pers.SourceSystem = NewSourceSystem(ss)
	pers.TiaEntityType = NewTiaEntityType("person")

	SourceID, err := NewSourceID(ss, ssID)
	if err != nil {
		return nil, err
	}
	pers.ID = SourceID.ID
	builder := PersonBuilder{pers: pers}
	return &builder, nil
}

// Name adds a name
func (pb *PersonBuilder) Name(n Name) *PersonBuilder {
	if pb.pers.Name == nil {
		pb.pers.Name = make([]Name, 0)
	}
	pb.pers.Name = append(pb.pers.Name, n)
	return pb
}

// ContactInfo adds contact information
func (pb *PersonBuilder) ContactInfo(ci ContactInfo) *PersonBuilder {
	if pb.pers.ContactInfo == nil {
		pb.pers.ContactInfo = make([]ContactInfo, 0)
	}
	pb.pers.ContactInfo = append(pb.pers.ContactInfo, ci)

	return pb
}

// Car adds a car
func (pb *PersonBuilder) Car(c Car) *PersonBuilder {
	if pb.pers.Cars == nil {
		pb.pers.Cars = make([]Car, 0)
	}
	pb.pers.Cars = append(pb.pers.Cars, c)
	return pb
}

// Status adds a status
func (pb *PersonBuilder) Status(s Status) *PersonBuilder {
	if pb.pers.Status == nil {
		pb.pers.Status = make([]Status, 0)
	}
	pb.pers.Status = append(pb.pers.Status, s)
	return pb
}

// DateOfBirth sets date of birth
func (pb *PersonBuilder) DateOfBirth(dob time.Time) *PersonBuilder {
	if !dob.IsZero() {
		pb.pers.DateOfBirth = &LocalDate{Time: dob}
	}
	return pb
}

// Gender sets the gender
func (pb *PersonBuilder) Gender(g string) *PersonBuilder {
	if len(g) > 0 {
		pb.pers.Gender = g
	}
	return pb
}

// PreferredLanguage adds a preferred language
func (pb *PersonBuilder) PreferredLanguage(lang string) *PersonBuilder {
	if len(lang) > 0 {
		pb.pers.PreferredLanguage = lang
	}
	return pb
}

// NativeTongue adds native tongue
func (pb *PersonBuilder) NativeTongue(nt string) *PersonBuilder {
	if len(nt) > 0 {
		if pb.pers.NativeTongue == nil {
			pb.pers.NativeTongue = make([]string, 0)
		}
		pb.pers.NativeTongue = append(pb.pers.NativeTongue, nt)
	}
	return pb
}

// SpecialNeed adds special need
func (pb *PersonBuilder) SpecialNeed(sn string) *PersonBuilder {
	if len(sn) > 0 {
		if pb.pers.SpecialNeed == nil {
			pb.pers.SpecialNeed = make([]string, 0)
		}
		pb.pers.SpecialNeed = append(pb.pers.SpecialNeed, sn)
	}
	return pb
}

// EntityKey add an entity-key
func (pb *PersonBuilder) EntityKey(typ string, value string) *PersonBuilder {
	if len(typ) > 0 && len(value) > 0 {
		if pb.pers.EntityKeys == nil {
			pb.pers.EntityKeys = make([]EntityKey, 0)
		}
		pb.pers.EntityKeys = append(pb.pers.EntityKeys, MakeEntityKey(typ, value))
	}
	return pb
}

// Role adds a new role
func (pb *PersonBuilder) Role(r Role) *PersonBuilder {
	if pb.pers.RoleList == nil {
		pb.pers.RoleList = make([]Role, 0)
	}
	pb.pers.RoleList = append(pb.pers.RoleList, r)
	return pb
}

// Build returns a new person struct
func (pb *PersonBuilder) Build() *Person {
	return pb.pers
}
