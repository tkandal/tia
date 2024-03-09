package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type TOPdeskBranch struct {
	Entity
	Name          string `json:"name" bson:"name"`
	Specification string `json:"specification,omitempty" bson:"specification,omitempty"`
	BranchType    string `json:"branchType,omitempty" bson:"branchType,omitempty"`
}
