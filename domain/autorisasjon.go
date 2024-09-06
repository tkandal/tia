package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Autorisasjon struct {
	Entity
}

type AutorisasjonBuilder struct {
	auth *Autorisasjon
}

func NewAutorisasjonBuilder(ss string, ssID string) (*AutorisasjonBuilder, error) {
	auth := &Autorisasjon{}
	auth.SourceSystem = NewSourceSystem(ss)
	auth.TiaEntityType = NewTiaEntityType("autorisasjon")

	sourceID, err := NewSourceID(ss, ssID)
	if err != nil {
		return nil, err
	}
	auth.ID = sourceID.ID
	builder := AutorisasjonBuilder{auth: auth}
	return &builder, nil
}

func (b *AutorisasjonBuilder) EntityKey(typ string, value string) *AutorisasjonBuilder {
	if len(typ) > 0 && len(value) > 0 {
		if b.auth.EntityKeys == nil {
			b.auth.EntityKeys = make([]EntityKey, 0)
		}
		b.auth.EntityKeys = append(b.auth.EntityKeys, MakeEntityKey(typ, value))
	}
	return b
}

func (b *AutorisasjonBuilder) Deleted() *AutorisasjonBuilder {
	t := time.Now()
	b.auth.Deleted = &t
	return b
}

func (b *AutorisasjonBuilder) Role(r Role) *AutorisasjonBuilder {
	if b.auth.RoleList == nil {
		b.auth.RoleList = make([]Role, 0)
	}
	b.auth.RoleList = append(b.auth.RoleList, r)
	return b
}

func (b *AutorisasjonBuilder) Build() *Autorisasjon {
	return b.auth
}
