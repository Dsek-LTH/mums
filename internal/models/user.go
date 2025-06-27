package models

type User struct {
	*Credentials
	Name string
}

func NewUser(credentials *Credentials, name string) *User {
	return &User{
		Credentials: credentials,
		Name: name,
	}
} 

