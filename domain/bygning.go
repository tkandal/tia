package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Bygning struct {
	Entity
	ID             int64  `json:"id,omitempty" bson:"id,omitempty"`
	BygningsNavn   string `json:"bygningsNavn,omitempty" bson:"bygningsNavn,omitempty"`
	Bygningsnummer string `json:"bygningsnummer,omitempty" bson:"bygningsnummer,omitempty"`
	GateAdresse    string `json:"gateAdresse,omitempty" bson:"gateAdresse,omitempty"`
	URL            string `json:"url,omitempty" bson:"url,omitempty"`
	URLNynorsk     string `json:"urlNynorsk,omitempty" bson:"urlNynorsk,omitempty"`
	URLEngelsk     string `json:"urlEngelsk,omitempty" bson:"urlEngelsk,omitempty"`
	DrawingsURL    string `json:"drawingsUrl,omitempty" bson:"drawingsUrl,omitempty"`
	PostalNum      string `json:"postalNum,omitempty" bson:"postalNum,omitempty"`
	PostalName     string `json:"postalName,omitempty" bson:"postalName,omitempty"`
}
