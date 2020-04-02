package model

import "time"

type Profile struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// profiles type profile
type Profiles []Profile

// NewProfile pada profiles constructor
func NewProfile() *Profile {
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
