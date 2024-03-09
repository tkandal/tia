package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type AuthInfo struct {
	TiaPersistable
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}
