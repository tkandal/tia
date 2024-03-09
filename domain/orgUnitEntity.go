package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
// Deprecated: These objects has been replaced, and should not be used.
type OrgUnitEntity struct {
	Entity
	OuID       int        `json:"ouId" bson:"ouId"`
	NsdCodes   []string   `json:"nsdCodes,omitempty" bson:"nsdCodes,omitempty"`
	ValidFrom  *time.Time `json:"validFrom,omitempty" bson:"validFrom,omitempty"`
	Expires    *time.Time `json:"expires,omitempty" bson:"expires,omitempty"`
	OuName     string     `json:"ouName,omitempty" bson:"ouName,omitempty"`
	OuAcronym  string     `json:"ouAcronym,omitempty" bson:"ouAcronym,omitempty"`
	OuCategory string     `json:"ouCategory,omitempty" bson:"ouCategory,omitempty"`
	OuEmail    string     `json:"ouEmail,omitempty" bson:"ouEmail,omitempty"`
	InstNr     int        `json:"instNr,omitempty" bson:"instNr,omitempty"`
	FacNr      int        `json:"facNr,omitempty" bson:"facNr,omitempty"`
	SecNr      int        `json:"secNr,omitempty" bson:"secNr,omitempty"`
	GrpNr      int        `json:"grpNr,omitempty" bson:"grpNr,omitempty"`
}
