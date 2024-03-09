package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Registerkort struct {
	TiaPersistable
	Arstall            string     `json:"arstall,omitempty" bson:"arstall,omitempty"`
	Terminkode         string     `json:"terminkode,omitempty" bson:"terminkode,omitempty"`
	StatusBetOk        string     `json:"statusBetOk,omitempty" bson:"statusBetOk,omitempty"`
	DatoBetaling       *time.Time `json:"datoBetaling,omitempty" bson:"datoBetaling,omitempty"`
	StatusRegOk        string     `json:"statusRegOk,omitempty" bson:"statusRegOk,omitempty"`
	DatoEndring        *time.Time `json:"datoEndring,omitempty" bson:"datoEndring,omitempty"`
	InstitusjonsNrEier int        `json:"institusjonsNrEier,omitempty" bson:"institusjonsNrEier,omitempty"`
	CampusKode         string     `json:"campusKode,omitempty" bson:"campusKode,omitempty"`
}
