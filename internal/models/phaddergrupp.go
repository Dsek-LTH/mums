package models

type Phaddergrupp struct {
	PrimaryColor string
	SecondaryColor string
	Icon string
}

func NewPhaddergrupp(primaryColor string, secondaryColor string, icon string) *Phaddergrupp {
	return &Phaddergrupp{
		PrimaryColor: primaryColor,
		SecondaryColor: secondaryColor,
		Icon: icon,
	}
}

