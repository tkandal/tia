package domain

import (
	"encoding/json"
	"strings"
)

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type PersonStudent struct {
	Person
	RegisterkortListe []Registerkort `json:"registerkortListe,omitempty" bson:"registerkortListe,omitempty"`
}

func (ps *PersonStudent) String() string {
	buf := &strings.Builder{}
	if err := json.NewEncoder(buf).Encode(ps); err != nil {
		return ""
	}
	return buf.String()
}
