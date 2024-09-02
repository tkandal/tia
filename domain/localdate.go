package domain

import (
	"encoding/json"
	"time"
)

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

const (
	dateFormatISO = "2006-01-02"
)

// LocalDate represents a date in ISO-format (YYYY-mm-dd).
type LocalDate struct {
	time.Time
}

// MarshalJSON encode this LocalDate-object to JSON.
func (ld *LocalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(ld.Format(dateFormatISO))
}

// UnmarshalJSON decodes JSON to a LocalDate-object.
func (ld *LocalDate) UnmarshalJSON(v []byte) error {
	var str string
	if err := json.Unmarshal(v, &str); err != nil {
		return err
	}
	t, err := time.Parse(dateFormatISO, str)
	if err != nil {
		return err
	}
	ld.Time = t
	return nil
}
