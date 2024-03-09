package domain

import (
	"fmt"
	"strings"
)

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// SourceID represents the combination of source-system and primary key for that source-system
type SourceID struct {
	ID             string `json:"id,omitempty" bson:"id,omitempty"`
	SourceSystem   string `json:"sourceSystem,omitempty" bson:"sourceSystem,omitempty"`
	SourceSystemID string `json:"sourceSystemId,omitempty" bson:"sourceSystemId,omitempty"`
}

// NewSourceID returns a pointer to a new SourceID struct
func NewSourceID(ss string, ssID string) (*SourceID, error) {
	if len(ss) == 0 {
		return nil, fmt.Errorf("source-system is empty")
	}
	if len(ssID) == 0 {
		return nil, fmt.Errorf("source-system identifier is empty")
	}
	if strings.HasPrefix(ssID, "urn:") {
		return nil, fmt.Errorf("source-system identifier is a full source-system urn %s", ssID)
	}
	sourceID := &SourceID{SourceSystem: ss, SourceSystemID: ssID}
	if ss[len(ss)-1] == '/' {
		ss = ss[0 : len(ss)-1]
	}
	sourceID.ID = fmt.Sprintf("%s/%s", ss, ssID)
	return sourceID, nil
}
