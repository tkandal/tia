package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */

type ArsTermin struct {
	TiaPersistable
	TerminKode        string `json:"terminKode,omitempty" bson:"terminKode,omitempty"`
	TerminNavn        string `json:"terminNavn,omitempty" bson:"terminNavn,omitempty"`
	TerminNavnNynorsk string `json:"terminNavnNynorsk,omitempty" bson:"terminNavnNynorsk,omitempty"`
	TerminNavnEngelsk string `json:"terminNavnEngelsk,omitempty" bson:"terminNavnEngelsk,omitempty"`
	DagNrFra          int    `json:"dagNrFra,omitempty" bson:"dagNrFra,omitempty"`
	DagNrTil          int    `json:"dagNrTil,omitempty" bson:"dagNrTil,omitempty"`
	ManedNrFra        int    `json:"manedNrFra,omitempty" bson:"manedNrFra,omitempty"`
	ManedNrTil        int    `json:"manedNrTil,omitempty" bson:"manedNrTil,omitempty"`
	StatusAktiv       string `json:"statusAktiv,omitempty" bson:"statusAktiv,omitempty"`
}
