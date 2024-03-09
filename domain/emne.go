package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
const (
	TypeNameShort = "forkortet"
	TypeName      = "navn"
)

type Emne struct {
	Entity
	Institusjonkode               int        `json:"institusjonkode,omitempty" bson:"institusjonkode,omitempty"`
	Emnekode                      string     `json:"emnekode,omitempty" bson:"emnekode,omitempty"`
	Versjonskode                  string     `json:"versjonskode,omitempty" bson:"versjonskode,omitempty"`
	Studiepoeng                   float32    `json:"studiepoeng,omitempty" bson:"studiepoeng,omitempty"`
	VektTypeKode                  string     `json:"vektTypeKode,omitempty" bson:"vektTypeKode,omitempty"`
	KarRegelkode                  int        `json:"karRegelkode,omitempty" bson:"karRegelkode,omitempty"`
	FagKodeSortering              string     `json:"fagKodeSortering,omitempty" bson:"fagKodeSortering,omitempty"`
	StudieprogramkodeRapportering string     `json:"studieprogramkodeRapportering,omitempty" bson:"studieprogramkode_rapportering,omitempty"`
	URL                           string     `json:"url,omitempty" bson:"url,omitempty"`
	EmneLopeNr                    int        `json:"emneLopeNr,omitempty" bson:"emneLopeNr,omitempty"`
	DatoSistEndret                *time.Time `json:"datoSistEndret,omitempty" bson:"datoSistEndret,omitempty"`
	LmsLopeNr                     int        `json:"lmsLopeNr,omitempty" bson:"lmsLopeNr,omitempty"`
	NameList                      []Name     `json:"nameList,omitempty" bson:"nameList,omitempty"`
}
