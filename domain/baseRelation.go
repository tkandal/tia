package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type BaseRelation struct {
	TiaPersistable
	Type            string `json:"type,omitempty" bson:"type,omitempty"`
	PrimaryRelation bool   `json:"primaryRelation" bson:"primaryRelation,omitempty"`
}
