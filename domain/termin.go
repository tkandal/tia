package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Termin struct {
	Entity
	TerminId    string     `json:"terminId" bson:"terminId"`
	Year        int        `json:"year" bson:"year"`
	Termin      string     `json:"termin" bson:"termin"`
	TerminStart *time.Time `json:"terminStart,omitempty" bson:"terminStart"`
	TerminEnd   *time.Time `json:"terminEnd,omitempty" bson:"terminEnd,omitempty"`
}
