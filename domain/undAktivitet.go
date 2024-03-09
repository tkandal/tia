package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type UndAktivitet struct {
	Entity
	Aktivitetkode    string `json:"aktivitetkode,omitempty" bson:"aktivitetkode,omitempty"`
	LmsRomMalKode    string `json:"lmsRomMalKode,omitempty" bson:"lmsRomMalKode,omitempty"`
	StatusEksportLms string `json:"statusEksportLms,omitempty" bson:"statusEksportLms,omitempty"`
	EksportLms       bool   `json:"eksportLms" bson:"eksportLms"`
	UndFormKode      string `json:"undFormKode,omitempty" bson:"undFormKode,omitempty"`
	UndPartiLopeNr   int    `json:"undPartiLopeNr,omitempty" bson:"undPartiLopeNr,omitempty"`
	DisiplinKode     string `json:"disiplinKode,omitempty" bson:"disiplinKode,omitempty"`
	Navn             string `json:"navn" bson:"navn"`
	Campus           string `json:"campus,omitempty" bson:"campus,omitempty"`
}
