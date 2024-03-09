package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// SourceSystem represents a source-system in TIA
type SourceSystem struct {
	TiaPersistable
	SystemUrn string `json:"systemUrn,omitempty" bson:"systemUrn,omitempty"`
}

// MakeSourceSystem makes a new SourceSystem struct
func MakeSourceSystem(urn string) SourceSystem {
	return SourceSystem{
		SystemUrn: urn,
	}
}

// NewSourceSystem return a pointer to a new SourceSystem struct
func NewSourceSystem(urn string) *SourceSystem {
	return &SourceSystem{
		SystemUrn: urn,
	}
}
