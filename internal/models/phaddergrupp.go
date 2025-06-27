package models

type Phaddergrupp struct {
	SwishNumber string
	PrimaryColor string
	SecondaryColor string
	Icon string
}

func NewPhaddergrupp(swishNumber string, primaryColor string, secondaryColor string, icon string) *Phaddergrupp {
	return &Phaddergrupp{
		SwishNumber: swishNumber,
		PrimaryColor: primaryColor,
		SecondaryColor: secondaryColor,
		Icon: icon,
	}
}

