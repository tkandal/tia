package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type OrgPosition struct {
	TiaPersistable
	Position string `json:"position" bson:"position"`
}
