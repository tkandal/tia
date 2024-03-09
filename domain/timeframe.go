package domain

import (
	"time"
)

/*
 * Copyright (c) 2019 Norwegian University of Science and Technology
 */

// Timeframe holds the dates for which an entity is active
type Timeframe struct {
	TiaPersistable
	From *time.Time `json:"from,omitempty" bson:"from,omitempty"`
	To   *time.Time `json:"to,omitempty" bson:"to,omitempty"`
}

// MakeTimeframe makes a new Timeframe struct
func MakeTimeframe(f time.Time, t time.Time) Timeframe {
	tf := Timeframe{}
	if !f.IsZero() {
		tf.From = &f
	}
	if !t.IsZero() {
		tf.To = &t
	}
	return tf
}

// NewTimeframe return a pointer a new Timeframe struct
func NewTimeframe(f time.Time, t time.Time) *Timeframe {
	tf := MakeTimeframe(f, t)
	return &tf
}
