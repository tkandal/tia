package domain

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// Privacy represents reservations in TIA
type Privacy struct {
	TiaSuper
	DoNotDistributeObject bool `json:"doNotDistributeObject,omitempty" bson:"doNotDistributeObject,omitempty"`
	DoNotShowObject       bool `json:"doNotShowObject,omitempty" bson:"doNotShowObject,omitempty"`
}
