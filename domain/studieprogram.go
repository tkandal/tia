package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Studieprogram struct {
	Entity
	StudieprogramKode string  `json:"studieprogramKode" bson:"studieprogramKode"`
	NameList          []Name  `json:"nameList,omitempty" bson:"nameList,omitempty"`
	StudieansvarSted  string  `json:"studieansvarSted,omitempty" bson:"studieansvarSted,omitempty"`
	Utdomradekode     string  `json:"utdomradekode,omitempty" bson:"utdomradekode,omitempty"`
	Vekttypekode      string  `json:"vekttypekode,omitempty" bson:"vekttypekode,omitempty"`
	URL               string  `json:"url,omitempty" bson:"url,omitempty"`
	Vektingstall      float32 `json:"vektingstall,omitempty" bson:"vektingstall,omitempty"`
	StudieNivaKode    int     `json:"studieNivaKode,omitempty" bson:"studieNivaKode,omitempty"`
}
