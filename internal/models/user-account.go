package models

type UserAccount struct {
	UserCredentials *UserCredentials
	UserProfile *UserProfile
}

func NewUserAccount(userCredentials *UserCredentials, userProfile *UserProfile) *UserAccount {
	return &UserAccount{
		UserCredentials: userCredentials,
		UserProfile: userProfile,
	}	
}

