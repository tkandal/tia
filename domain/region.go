package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Region struct {
	Entity
	ID         string `json:"id" bson:"id"`
	RegionNavn string `json:"regionNavn,omitempty" bson:"regionNavn,omitempty"`
}
