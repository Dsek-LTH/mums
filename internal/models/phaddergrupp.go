package models

import "github.com/Dsek-LTH/mums/internal/config"

type Phaddergrupp struct {
	Name string
	IconFilePath string
	PrimaryColor string
	SecondaryColor string
	SwishNumber string
	MumsPrice int64
	PaymentMessage string
}

func NewPhaddergrupp(name string, iconFilePath string) *Phaddergrupp {
	return &Phaddergrupp{
		Name: name,
		IconFilePath: iconFilePath,
		PrimaryColor: config.DefaultPrimaryPhaddergruppColor,
		SecondaryColor: config.DefaultSecondaryPhaddergruppColor,
		SwishNumber: "",
		MumsPrice: 10,
		PaymentMessage: config.DefaultPaymentMessage,
	}
}

