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

type OrganisasjonBuilder struct {
	org *Organisasjon
}

func NewOrganisasjonBuilder(ss string, ssID string) (*OrganisasjonBuilder, error) {
	org := &Organisasjon{}
	org.SourceSystem = NewSourceSystem(ss)
	org.TiaEntityType = NewTiaEntityType("organisasjon")

	sourceID, err := NewSourceID(ss, ssID)
	if err != nil {
		return nil, err
	}
	org.ID = sourceID.ID
	builder := &OrganisasjonBuilder{org: org}
	return builder, nil
}

func (ob *OrganisasjonBuilder) Name(name Name) *OrganisasjonBuilder {
	if ob.org.Name == nil {
		ob.org.Name = make([]Name, 0)
	}
	ob.org.Name = append(ob.org.Name, name)
	return ob
}

func (ob *OrganisasjonBuilder) ContactInfo(ci ContactInfo) *OrganisasjonBuilder {
	if ob.org.ContactInfo == nil {
		ob.org.ContactInfo = make([]ContactInfo, 0)
	}
	ob.org.ContactInfo = append(ob.org.ContactInfo, ci)
	return ob
}

func (ob *OrganisasjonBuilder) Status(status Status) *OrganisasjonBuilder {
	if ob.org.Status == nil {
		ob.org.Status = make([]Status, 0)
	}
	ob.org.Status = append(ob.org.Status, status)
	return ob
}

func (ob *OrganisasjonBuilder) Description(desc string) *OrganisasjonBuilder {
	if len(desc) > 0 {
		ob.org.Description = desc
	}
	return ob
}

func (ob *OrganisasjonBuilder) ValidFrom(t *time.Time) *OrganisasjonBuilder {
	if t != nil && !t.IsZero() {
		ob.org.ValidFrom = t
	}
	return ob
}

func (ob *OrganisasjonBuilder) ValidTo(t *time.Time) *OrganisasjonBuilder {
	if t != nil && !t.IsZero() {
		ob.org.ValidTo = t
	}
	return ob
}

func (ob *OrganisasjonBuilder) EntityKey(typ string, value string) *OrganisasjonBuilder {
	if len(typ) > 0 && len(value) > 0 {
		if ob.org.EntityKeys == nil {
			ob.org.EntityKeys = make([]EntityKey, 0)
		}
		ob.org.EntityKeys = append(ob.org.EntityKeys, MakeEntityKey(typ, value))
	}
	return ob
}

func (ob *OrganisasjonBuilder) Role(r Role) *OrganisasjonBuilder {
	if ob.org.RoleList == nil {
		ob.org.RoleList = make([]Role, 0)
	}
	ob.org.RoleList = append(ob.org.RoleList, r)
	return ob
}

func (ob *OrganisasjonBuilder) Build() *Organisasjon {
	return ob.org
}
