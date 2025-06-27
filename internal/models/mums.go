package models

import "time"

type Mums struct {
	UserAccount *UserAccount
	Timestamp time.Time 
}

func NewMums(userAccount *UserAccount, timestamp time.Time) *Mums {
	return &Mums{
		UserAccount: userAccount,
		Timestamp: timestamp,
	}
}

