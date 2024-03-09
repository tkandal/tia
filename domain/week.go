package domain

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type Week struct {
	TiaPersistable
	WeekID     string `json:"weekId,omitempty" bson:"weekId,omitempty"`
	Year       int    `json:"year,omitempty" bson:"year,omitempty"`
	WeekNumber int    `json:"weekNumber,omitempty" bson:"weekNumber,omitempty"`
}
