package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type UndervisningsEnhet struct {
	Entity
	TerminNr          int        `json:"terminNr,omitempty" bson:"terminNr,omitempty"`
	LmsRomMalKode     string     `json:"lmsRomMalKode,omitempty" bson:"lmsRomMalKode,omitempty"`
	StatusEksportLms  string     `json:"statusEksportLms,omitempty" bson:"statusEksportLms,omitempty"`
	LmsLopeNr         int        `json:"lmsLopeNr,omitempty" bson:"lmsLopeNr,omitempty"`
	DatoEksportLmsFra *time.Time `json:"datoEksportLmsFra,omitempty" bson:"datoEksportLmsFra,omitempty"`
	DatoEksportLmsTil *time.Time `json:"datoEksportLmsTil,omitempty" bson:"datoEksportLmsTil,omitempty"`
	AntallTerminer    int        `json:"antallTerminer,omitempty" bson:"antallTerminer,omitempty"`
}
