package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Tjenesteforhold struct {
	Entity
	TjenForholdkode       string `json:"tjenForholdkode,omitempty" bson:"tjenForholdkode,omitempty"`
	TjenForhold           string `json:"tjenForhold,omitempty" bson:"tjenForhold,omitempty"`
	TjenForholdRangering  string `json:"tjenForholdRangering,omitempty" bson:"tjenForholdRangering,omitempty"`
	LonnsKode             string `json:"lonnsKode,omitempty" bson:"lonnsKode,omitempty"`
	LonnsKodetekst        string `json:"lonns_kodetekst,omitempty" bson:"lonns_kodetekst,omitempty"`
	LonnsKodeRangering    string `json:"lonnsKodeRangering,omitempty" bson:"lonnsKodeRangering,omitempty"`
	StatusBaseksport      string `json:"statusBaseksport,omitempty" bson:"statusBaseksport,omitempty"`
	BasRolle              string `json:"basRolle,omitempty" bson:"basRolle,omitempty"`
	StatusPersontjenesten string `json:"statusPersontjenesten,omitempty" bson:"statusPersontjenesten,omitempty"`
	Affiliation           string `json:"affiliation,omitempty" bson:"affiliation,omitempty"`
	StatusAlumnieksport   string `json:"statusAlumnieksport,omitempty" bson:"statusAlumnieksport,omitempty"`
	StatusFseksport       string `json:"statusFseksport,omitempty" bson:"statusFseksport,omitempty"`
	StatusLonneksport     string `json:"statusLonneksport,omitempty" bson:"statusLonneksport,omitempty"`
	StatusAdgangeksport   string `json:"statusAdgangeksport,omitempty" bson:"statusAdgangeksport,omitempty"`
	StatusProsjekteksport string `json:"statusProsjekteksport,omitempty" bson:"statusProsjekteksport,omitempty"`
	StatusLiseksport      string `json:"statusLiseksport,omitempty" bson:"statusLiseksport,omitempty"`
	StatusArkiveksport    string `json:"statusArkiveksport,omitempty" bson:"statusArkiveksport,omitempty"`
	StatusHmseksport      string `json:"statusHmseksport,omitempty" bson:"statusHmseksport,omitempty"`
	StatusStolaveksport   string `json:"statusStolaveksport,omitempty" bson:"statusStolaveksport,omitempty"`
	StatusCristineksport  string `json:"statusCristineksport,omitempty" bson:"statusCristineksport,omitempty"`
	StatusBibsyseksport   string `json:"statusBibsyseksport,omitempty" bson:"statusBibsyseksport,omitempty"`
	TjenForholdMerknad    string `json:"tjenForholdMerknad,omitempty" bson:"tjenForholdMerknad,omitempty"`
}
