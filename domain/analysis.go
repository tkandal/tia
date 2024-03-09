package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */

type Analysis struct {
	Entity
	AnId          string `json:"anId,omitempty" bson:"anId,omitempty"`
	ContactPerson string `json:"contactPerson,omitempty" bson:"contactPerson,omitempty"`
	Description   string `json:"description,omitempty" bson:"description,omitempty"`
}
