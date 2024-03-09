package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type VurderingPart struct {
	TiaPersistable
	VurderingsNavn   string     `json:"vurderingsNavn,omitempty" bson:"vurderingsNavn,omitempty"`
	Vurderingsform   string     `json:"vurderingsform,omitempty" bson:"vurderingsform,omitempty"`
	AnnonseringsDato *time.Time `json:"annonseringsDato,omitempty" bson:"annonseringsDato,omitempty"`
	StartTid         *time.Time `json:"startTid,omitempty" bson:"startTid,omitempty"`
	Varighet         int        `json:"varighet,omitempty" bson:"varighet,omitempty"`
	TidsEnhet        string     `json:"tidsEnhet,omitempty" bson:"tidsEnhet,omitempty"`
	Type             string     `json:"type,omitempty" bson:"type,omitempty"`
	Rom              string     `json:"rom,omitempty" bson:"rom,omitempty"`
	Bygning          string     `json:"bygning,omitempty" bson:"bygning,omitempty"`
}
