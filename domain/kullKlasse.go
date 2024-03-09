package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type KullKlasse struct {
	Entity
	StatusAktiv      string `json:"statusAktiv,omitempty" bson:"statusAktiv,omitempty"`
	StatusEksportLms string `json:"statusEksportLms,omitempty" bson:"statusEksportLms,omitempty"`
	LmsRomMalKode    string `json:"lmsRomMalKode,omitempty" bson:"lmsRomMalKode,omitempty"`
}
