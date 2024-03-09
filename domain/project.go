package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Project struct {
	Entity
	PrId              string `json:"prId" bson:"prId"`
	Description       string `json:"description,omitempty" bson:"description,omitempty"`
	Leader            string `json:"leader,omitempty" bson:"leader,omitempty"`
	Group             string `json:"group,omitempty" bson:"group,omitempty"`
	FinancingProgram  string `json:"financingProgram,omitempty" bson:"financingProgram,omitempty"`
	FinancingCategory string `json:"financingCategory,omitempty" bson:"financingCategory,omitempty"`
	EnabledFlag       bool   `json:"enabledFlag,omitempty" bson:"enabledFlag,omitempty"`
}
