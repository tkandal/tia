package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Etasje struct {
	Entity
	ID         int64  `json:"id,omitempty" bson:"id,omitempty"`
	EtasjeNavn string `json:"etasjeNavn,omitempty" bson:"etasjeNavn,omitempty"`
	EtasjeKode string `json:"etasjeKode,omitempty" bson:"etasjeKode,omitempty"`
}
