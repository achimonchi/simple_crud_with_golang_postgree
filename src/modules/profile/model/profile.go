package model

import "time"

// membuat struct untuk profile
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
// ini berguna untuk mendapatkan seluru data dari Profile
type Profiles []Profile

// NewProfile pada profile constructor
func NewProfile() *Profile {
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
