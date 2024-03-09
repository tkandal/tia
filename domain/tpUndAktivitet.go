package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type TPUndAktivitet struct {
	Entity
	Terminnr      int    `json:"terminnr,omitempty" bson:"terminnr,omitempty"`
	AktivitetKode string `json:"aktivitetKode,omitempty" bson:"aktivitetKode,omitempty"`
	ArsterminId   string `json:"arsterminId,omitempty" bson:"arsterminId,omitempty"`
	Metode        string `json:"metode,omitempty" bson:"metode,omitempty"`
	MetodeNavn    string `json:"metodeNavn,omitempty" bson:"metodeNavn,omitempty"`
	Tittel        string `json:"tittel,omitempty" bson:"tittel,omitempty"`
	TPID          string `json:"tpid" bson:"tpid"`
	CourseID      string `json:"courseId,omitempty" bson:"courseId,omitempty"`
}
