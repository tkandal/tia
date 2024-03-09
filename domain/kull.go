package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Kull struct {
	Entity
	StatusAutomatiskOpprykk string `json:"statusAutomatiskOpprykk,omitempty" bson:"statusAutomatiskOpprykk,omitempty"`
	StatusGenererEpost      string `json:"statusGenererEpost,omitempty" bson:"statusGenererEpost,omitempty"`
	StatusAktiv             string `json:"statusAktiv,omitempty" bson:"statusAktiv,omitempty"`
	StatusEksportLms        string `json:"statusEksportLms,omitempty" bson:"statusEksportLms,omitempty"`
	LmsRomMalKode           string `json:"lmsRomMalKode,omitempty" bson:"lmsRomMalKode,omitempty"`
	LmsLopeNr               int    `json:"lmsLopeNr,omitempty" bson:"lmsLopeNr,omitempty"`
	KulltrinnStart          int    `json:"kulltrinnStart,omitempty" bson:"kulltrinnStart,omitempty"`
	Bibsysstedkode          string `json:"bibsysstedkode,omitempty" bson:"bibsysstedkode,omitempty"`
	Arstall                 int    `json:"arstall,omitempty" bson:"arstall,omitempty"`
	Terminkode              string `json:"terminkode,omitempty" bson:"terminkode,omitempty"`
	Studieprogramkode       string `json:"studieprogramkode,omitempty" bson:"studieprogramkode,omitempty"`
	Studiekullnavn          string `json:"studiekullnavn,omitempty" bson:"studiekullnavn,omitempty"`
	StudiekullnavnEngelsk   string `json:"studiekullnavnEngelsk,omitempty" bson:"studiekullnavnEngelsk,omitempty"`
}
