package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type TPStaff struct {
	Fornavn   string `json:"fornavn" bson:"fornavn"`
	Etternavn string `json:"etternavn" bson:"etternavn"`
	Kortnavn  string `json:"kortnavn,omitempty" bson:"kortnavn,omitempty"`
	URL       string `json:"url,omitempty" bson:"url,omitempty"`
}
