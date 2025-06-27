package models

type UserProfile struct {
	Name string
}

func NewUserProfile(name string) *UserProfile {
	return &UserProfile{
		Name: name,
	}	
}

