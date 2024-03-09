package domain

import (
	"time"
)

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// Entity is a struct that all actual data-elements inherit in TIA
type Entity struct {
	TiaPersistable
	ID                  string         `json:"_id,omitempty" bson:"id,omitempty"`
	Checksum            string         `json:"_checksum,omitempty" bson:"_checksum,omitempty"`
	SourceSystem        *SourceSystem  `json:"sourceSystem,omitempty" bson:"sourceSystem,omitempty"`
	SourceSystemUpdated int64          `json:"sourceSystemUpdated,omitempty" bson:"sourceSystemUpdated,omitempty"`
	Deleted             *time.Time     `json:"deleted,omitempty" bson:"deleted,omitempty"`
	EntityKeys          []EntityKey    `json:"entityKeys,omitempty" bson:"entityKeys,omitempty"`
	TiaEntityType       *TiaEntityType `json:"entityType,omitempty" bson:"entityType,omitempty"`
	RoleList            []Role         `json:"roleList,omitempty" bson:"roleList,omitempty"`
}
