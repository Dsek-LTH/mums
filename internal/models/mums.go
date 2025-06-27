package models

import "time"

type Mums struct {
	User *User
	Timestamp time.Time 
}

func NewMums(user *User, timestamp time.Time) *Mums {
	return &Mums{
		User: user,
		Timestamp: timestamp,
	}
}

