package models

type PhaddergruppMapping struct {
	UserAccount *UserAccount
	Phaddergrupp *Phaddergrupp
	PhaddergruppRole PhaddergruppRole
	MumsAvailable int16
}

func NewPhaddergruppMapping(userAccount *UserAccount, phaddergrupp *Phaddergrupp, phaddergruppRole PhaddergruppRole) *PhaddergruppMapping {
	return &PhaddergruppMapping{
		UserAccount: userAccount,
		Phaddergrupp: phaddergrupp,
		PhaddergruppRole: phaddergruppRole,
		MumsAvailable: 0,
	}
}

