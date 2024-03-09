package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Campus struct {
	Entity
	CampusNavn        string     `json:"campusNavn,omitempty" bson:"campusNavn,omitempty"`
	CampusNavnNynorsk string     `json:"campusNavnNynorsk,omitempty" bson:"campusNavnNynorsk,omitempty"`
	CampusNavnEngelsk string     `json:"campusNavnEngelsk,omitempty" bson:"campusNavnEngelsk,omitempty"`
	LandNr            int        `json:"landNr,omitempty" bson:"landNr,omitempty"`
	FylkesNr          int        `json:"fylkesNr,omitempty" bson:"fylkesNr,omitempty"`
	KommuneNr         int        `json:"kommuneNr,omitempty" bson:"kommuneNr,omitempty"`
	DatoOpprettet     *time.Time `json:"datoOpprettet,omitempty" bson:"datoOpprettet,omitempty"`
	DatoEndret        *time.Time `json:"datoEndret,omitempty" bson:"datoEndret,omitempty"`
	Orgnr             string     `json:"orgnr,omitempty" bson:"orgnr,omitempty"`
	StatusDefault     string     `json:"statusDefault,omitempty" bson:"statusDefault,omitempty"`
}
