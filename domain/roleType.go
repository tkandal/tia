package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// RoleType represents a role-type in TIA
type RoleType struct {
	Entity
	RoleCode string `json:"roleCode,omitempty" bson:"roleCode,omitempty"`
}

// RoleTypeBuilder is a role-type builder struct
type RoleTypeBuilder struct {
	roleType RoleType
}

// NewRoleTypeBuilder return a pointer to a role-type builder
func NewRoleTypeBuilder() *RoleTypeBuilder {
	rt := RoleType{}
	rt.TiaEntityType = NewTiaEntityType("rolletype")
	return &RoleTypeBuilder{roleType: rt}
}

// RoleCode sets the type of role
func (rb *RoleTypeBuilder) RoleCode(rc string) *RoleTypeBuilder {
	if len(rc) > 0 {
		rb.roleType.RoleCode = rc
		rb.roleType.EntityKeys = []EntityKey{MakeEntityKey(IdentityTypeRoleTypeKeyname, rc)}
	}
	return rb
}

// Build return a RoleType
func (rb *RoleTypeBuilder) Build() RoleType {
	return rb.roleType
}
