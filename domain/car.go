package domain

import (
	"time"
)

/*
 * Copyright (c) 2019, 2022 Norwegian University of Science and Technology
 */

// Car represents a car in TIA.
// swagger:model Car
type Car struct {
	TiaPersistable
	RegistrationTime *time.Time `json:"registrationTime,omitempty" bson:"registrationTime,omitempty"`
	Plate            string     `json:"plate,omitempty" bson:"plate,omitempty"`
	Permissions      []string   `json:"permissions,omitempty" bson:"permissions,omitempty"`
}

// CarBuilder is a struct that builds a care
type CarBuilder struct {
	car Car
}

// NewCarBuilder starts a new car-builder
func NewCarBuilder(plate string) *CarBuilder {
	c := Car{}
	if len(plate) > 0 {
		c.Plate = plate
	}
	return &CarBuilder{car: c}
}

// RegistrationTime register the time when a car was imported to TIA
func (cb *CarBuilder) RegistrationTime(rt time.Time) *CarBuilder {
	if !rt.IsZero() {
		cb.car.RegistrationTime = &rt
	}
	return cb
}

// Permissions is permission for parking at NTNU
func (cb *CarBuilder) Permissions(perms []string) *CarBuilder {
	if cb.car.Permissions == nil {
		cb.car.Permissions = []string{}
	}
	cb.car.Permissions = append(cb.car.Permissions, perms...)
	return cb
}

// Build finally builds the car
func (cb *CarBuilder) Build() Car {
	return cb.car
}
