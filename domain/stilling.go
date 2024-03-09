package domain

import (
	"git.it.ntnu.no/df/tia/backend/appapi-push/types"
)

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Stilling struct {
	TiaPersistable
	Hovedarbeidsforhold          bool             `json:"hovedarbeidsforhold" bson:"hovedarbeidsforhold"`
	TjenStedNr                   string           `json:"tjenStedNr,omitempty" bson:"tjenStedNr,omitempty"`
	AnsvStedNr                   string           `json:"ansvStedNr,omitempty" bson:"ansvStedNr,omitempty"`
	KostStedNr                   string           `json:"kostStedNr,omitempty" bson:"kostStedNr"`
	StillBetegnelse              string           `json:"stillBetegnelse,omitempty" bson:"stillBetegnelse,omitempty"`
	StillKode                    string           `json:"stillKode,omitempty" bson:"stillKode,omitempty"`
	StillingsBeskrivelse         string           `json:"stillingsBeskrivelse,omitempty" bson:"stillingsBeskrivelse,omitempty"`
	Univkat                      string           `json:"univkat,omitempty" bson:"univkat,omitempty"`
	StillAndel                   float64          `json:"stillAndel,omitempty" bson:"stillAndel,omitempty"`
	StartNaverendeArbeidsforhold *types.LocalDate `json:"startNaverendeArbeidsforhold,omitempty" bson:"startNaverendeArbeidsforhold,omitempty"`
	FraDato                      *types.LocalDate `json:"fraDato,omitempty" bson:"fraDato,omitempty"`
	TilDato                      *types.LocalDate `json:"tilDato,omitempty" bson:"tilDato,omitempty"`
	FLonnsdag                    *types.LocalDate `json:"fLonnsdag,omitempty" bson:"fLonnsdag,omitempty"`
	SLonnsdag                    *types.LocalDate `json:"sLonnsdag,omitempty" bson:"sLonnsdag,omitempty"`
	LonnsTrinnA                  string           `json:"lonnsTrinnA,omitempty" bson:"lonnsTrinnA,omitempty"`
	LonnsTrinnB                  string           `json:"lonnsTrinnB,omitempty" bson:"lonnsTrinnB,omitempty"`
	BruttoMndLonn                float64          `json:"bruttoMndLonn,omitempty" bson:"bruttoMndLonn,omitempty"`
	TjenForholdNr                string           `json:"tjenForholdNr,omitempty" bson:"tjenForholdNr,omitempty"`
	AntTjenForhold               string           `json:"antTjenForhold,omitempty" bson:"antTjenForhold,omitempty"`
	Lønnsandel                   float64          `json:"lønnsandel,omitempty" bson:"lønnsandel,omitempty"`
	// keys
	TjenStedKey    *EntityKey `json:"tjenStedKey,omitempty" bson:"tjenStedKey,omitempty"`
	AnsvStedKey    *EntityKey `json:"ansvStedKey,omitempty" bson:"ansvStedKey,omitempty"`
	KostStedKey    *EntityKey `json:"kostStedKey,omitempty" bson:"kostStedKey,omitempty"`
	TjenForholdKey *EntityKey `json:"tjenForholdKey,omitempty" bson:"tjenForholdKey,omitempty"`
	LederSignKey   *EntityKey `json:"lederSignKey,omitempty" bson:"lederSignKey,omitempty"`
}
