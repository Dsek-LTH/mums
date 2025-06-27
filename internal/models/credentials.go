package models

type Credentials struct {
	Email string
	Hashword string
}

func NewCredentials(email string, hashword string) *Credentials {
	return &Credentials{
		Email: email,
		Hashword: hashword,
	}	
}

