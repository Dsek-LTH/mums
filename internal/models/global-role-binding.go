package models

type GlobalRoleBinding struct {
	User *User
	GlobalRole GlobalRole
}

