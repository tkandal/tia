package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Infotype struct {
	TiaPersistable
	Infotypekode        string `json:"infotypekode,omitempty" bson:"infotypekode,omitempty"`
	InfotypeNavn        string `json:"infotypeNavn,omitempty" bson:"infotypeNavn,omitempty"`
	StatusDefault       string `json:"statusDefault,omitempty" bson:"statusDefault,omitempty"`
	KodeverdiType       string `json:"kodeverdiType,omitempty" bson:"kodeverdiType,omitempty"`
	TagKode             string `json:"tagKode,omitempty" bson:"tagKode,omitempty"`
	StatusStudieprogram string `json:"statusStudieprogram,omitempty" bson:"statusStudieprogram,omitempty"`
	StatusEmne          string `json:"statusEmne,omitempty" bson:"statusEmne,omitempty"`
	InfotypeNavnEngelsk string `json:"infotypeNavnEngelsk,omitempty" bson:"infotypeNavnEngelsk,omitempty"`
	InfotypeNavnNynorsk string `json:"infotypeNavnNynorsk,omitempty" bson:"infotypeNavnNynorsk,omitempty"`
	CdmElementKode      string `json:"cdmElementKode,omitempty" bson:"cdmElementKode,omitempty"`
	RekkefolgeNr        int    `json:"rekkefolgeNr,omitempty" bson:"rekkefolgeNr,omitempty"`
	StatusKurs          string `json:"statusKurs,omitempty" bson:"statusKurs,omitempty"`
	StatusPubWebApp     string `json:"statusPubWebApp,omitempty" bson:"statusPubWebApp,omitempty"`
	StatusPubWeb        string `json:"statusPubWeb,omitempty" bson:"statusPubWeb,omitempty"`
	StatusUtveksl       string `json:"statusUtveksl,omitempty" bson:"statusUtveksl,omitempty"`
	InstitusjonsNrEier  int    `json:"institusjonsNrEier" bson:"institusjonsNrEier,omitempty"`
}
