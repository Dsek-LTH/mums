package models

type PhaddergruppBinding struct {
	User *User
	Phaddergrupp *Phaddergrupp
	PhaddergruppRole PhaddergruppRole
	MumsAvailable int16
}

func NewPhaddergruppBinding(user *User, phaddergrupp *Phaddergrupp, phaddergruppRole PhaddergruppRole) *PhaddergruppBinding {
	return &PhaddergruppBinding{
		User: user,
		Phaddergrupp: phaddergrupp,
		PhaddergruppRole: phaddergruppRole,
		MumsAvailable: 0,
	}
}

