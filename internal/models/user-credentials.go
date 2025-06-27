package models

type UserCredentials struct {
	Email string
	Hashword string
}

func NewUserCredentials(email string, hashword string) *UserCredentials {
	return &UserCredentials{
		Email: email,
		Hashword: hashword,
	}	
}

