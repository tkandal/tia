package domain

import "time"

/*
 * Copyright (c) 2021 Norwegian University of Science and Technology
 */
type UndervisningsEnhetBB struct {
	Entity
	BBCourseID   string     `json:"bbCourseId,omitempty" bson:"bbCourseId,omitempty"`
	UUID         string     `json:"uuid,omitempty" bson:"uuid,omitempty"`
	ExternalID   string     `json:"externalId,omitempty" bson:"externalId,omitempty"`
	DataSourceID string     `json:"dataSourceId,omitempty" bson:"dataSourceId,omitempty"`
	CourseID     string     `json:"courseId,omitempty" bson:"courseId,omitempty"`
	Name         string     `json:"name,omitempty" bson:"name,omitempty"`
	Created      *time.Time `json:"created,omitempty" bson:"created,omitempty"`
	AllowGuests  bool       `json:"allowGuests" bson:"allowGuests"`
	ReadOnly     bool       `json:"readOnly" bson:"readOnly"`
	TermID       string     `json:"termId,omitempty" bson:"termId,omitempty"`
}
