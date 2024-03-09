package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type TPUndAktHendelse struct {
	Ukenr      int    `json:"ukenr" bson:"ukenr"`
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	Tema       string `json:"tema,omitempty" bson:"tema,omitempty"`
	Curriculum string `json:"curriculum,omitempty" bson:"curriculum,omitempty"`
	Tittel     string `json:"tittel" bson:"tittel"`
	ID         string `json:"id" bson:"id"`
}
