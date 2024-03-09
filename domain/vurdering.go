package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Vurdering struct {
	Entity
	InstitusjonsNr    int        `json:"institusjonsNr,omitempty" bson:"institusjonsNr,omitempty"`
	Emnekode          string     `json:"emnekode,omitempty" bson:"emnekode,omitempty"`
	Versjonskode      string     `json:"versjonskode,omitempty" bson:"versjonskode,omitempty"`
	VurderingsTermin  string     `json:"vurderingsTermin,omitempty" bson:"vurderingsTermin,omitempty"`
	VurderingsAr      int        `json:"vurderingsAr,omitempty" bson:"vurderingsAr,omitempty"`
	VurderingsType    string     `json:"vurderingsType,omitempty" bson:"vurderingsType,omitempty"`
	RegistreringStart *time.Time `json:"registreringStart,omitempty" bson:"registreringStart,omitempty"`
	RegistreringSlutt *time.Time `json:"registreringSlutt,omitempty" bson:"registreringSlutt,omitempty"`
	StatusPublisering string     `json:"statusPublisering,omitempty" bson:"statusPublisering,omitempty"`
}
