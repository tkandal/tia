package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Rom struct {
	/* The system entity key for a (FS) room consists of a concatenated string with the
	   same parts that constitutes the compound primary key in the FS table "Rom":
	   Institusjonsnr_Eier_BYGNINGSKODE_ROMKODE
	*/

	// A concatenated unique index used in TIA Stage table: _bygningskode romkode
	SyllabusKey     string  `json:"syllabusKey" bson:"syllabusKey"`
	RomNavn         string  `json:"romNavn,omitempty" bson:"romNavn,omitempty"`
	Etasje          int     `json:"etasje,omitempty" bson:"etasje,omitempty"`
	Romtypekode     string  `json:"romtypekode,omitempty" bson:"romtypekode,omitempty"`
	KapasitetEks    int     `json:"kapasitetEks,omitempty" bson:"kapasitetEks,omitempty"`
	KapasitetUnd    int     `json:"kapasitetUnd,omitempty" bson:"kapasitetUnd,omitempty"`
	Merknad         string  `json:"merknad,omitempty" bson:"merknad,omitempty"`
	MerknadAnkomst  string  `json:"merknadAnkomst,omitempty" bson:"merknadAnkomst,omitempty"`
	URL             string  `json:"url,omitempty" bson:"url,omitempty"`
	StatusAktiv     string  `json:"statusAktiv,omitempty" bson:"statusAktiv,omitempty"`
	BygningsNavn    string  `json:"bygningsNavn,omitempty" bson:"bygningsNavn,omitempty"`
	BygningsAkronym string  `json:"bygningsAkronym,omitempty" bson:"bygningsAkronym,omitempty"`
	Romnummer       string  `json:"romnummer,omitempty" bson:"romnummer,omitempty"`
	RomKategori     string  `json:"romKategori,omitempty" bson:"romKategori,omitempty"`
	GrossArea       float64 `json:"grossArea,omitempty" bson:"grossArea,omitempty"`
	NetArea         float64 `json:"netArea,omitempty" bson:"netArea,omitempty"`
}
