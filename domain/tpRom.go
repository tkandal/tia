package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type TPRom struct {
	Romnavn     string `json:"romnavn" bson:"romnavn"`
	Bygningnavn string `json:"bygningnavn,omitempty" bson:"bygningnavn,omitempty"`
	RomURL      string `json:"romUrl,omitempty" bson:"romUrl,omitempty"`
}
