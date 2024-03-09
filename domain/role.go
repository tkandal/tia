package domain

import (
	"fmt"
	"time"
)

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// Role represents a role in TIA
type Role struct {
	TiaPersistable
	Timeframe  *Timeframe  `json:"timeframe,omitempty" bson:"timeframe,omitempty"`
	ActsOn     *EntityKey  `json:"actsOn,omitempty" bson:"actsOn,omitempty"`
	EntityKeys []EntityKey `json:"entityKeys,omitempty" bson:"entityKeys,omitempty"`
	Deleted    *time.Time  `json:"deleted,omitempty" bson:"deleted,omitempty"`
}

// RoleBuilder is a role-builder struct
type RoleBuilder struct {
	role Role
}

// NewRoleBuilder returns a pointer to a RoleBuilder struct
func NewRoleBuilder(typ string) (*RoleBuilder, error) {
	if len(typ) == 0 {
		return nil, fmt.Errorf("role-type is empty")
	}
	rt := NewRoleTypeBuilder().RoleCode(typ).Build()
	roleType := rt.EntityKeys[0]
	role := Role{}
	role.EntityKeys = []EntityKey{}
	role.EntityKeys = append(role.EntityKeys, roleType)
	return &RoleBuilder{role: role}, nil
}

// RoleType adds a role-type
func (rb *RoleBuilder) RoleType(rt RoleType) *RoleBuilder {
	if len(rt.EntityKeys) > 0 {
		entKey := rt.EntityKeys[0]
		rb.role.EntityKeys = append(rb.role.EntityKeys, entKey)
	}
	return rb
}

// Timeframe sets the time-frame
func (rb *RoleBuilder) Timeframe(from time.Time, to time.Time) *RoleBuilder {
	if !from.IsZero() {
		if rb.role.Timeframe == nil {
			rb.role.Timeframe = &Timeframe{}
		}
		rb.role.Timeframe.From = &from
	}
	if !to.IsZero() {
		if rb.role.Timeframe == nil {
			rb.role.Timeframe = &Timeframe{}
		}
		rb.role.Timeframe.To = &to
	}
	return rb
}

// ActsOn sets entity.keys for which this role acts on
func (rb *RoleBuilder) ActsOn(typ string, value string) *RoleBuilder {
	if len(typ) > 0 && len(value) > 0 {
		rb.role.ActsOn = NewEntityKey(typ, value)
	}
	return rb
}

// EntityKey adds entity-key
func (rb *RoleBuilder) EntityKey(typ string, value string) *RoleBuilder {
	if len(typ) > 0 && len(value) > 0 {
		if rb.role.EntityKeys == nil {
			rb.role.EntityKeys = []EntityKey{}
		}
		rb.role.EntityKeys = append(rb.role.EntityKeys, MakeEntityKey(typ, value))
	}
	return rb
}

// Deleted sets the time when this role is deleted
func (rb *RoleBuilder) Deleted(t time.Time) *RoleBuilder {
	if !t.IsZero() {
		rb.role.Deleted = &t
	}
	return rb
}

// Build returns the role
func (rb *RoleBuilder) Build() Role {
	return rb.role
}
