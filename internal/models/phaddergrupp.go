package models

import "github.com/Dsek-LTH/mums/internal/config"

type Phaddergrupp struct {
	Name string
	IconFilePath string
	PrimaryColor string
	SecondaryColor string
	MumsPrice int64
	SwishRecipientNumber string
	SwishRecipientName string
}

func NewPhaddergrupp(name string, iconFilePath string) *Phaddergrupp {
	return &Phaddergrupp{
		Name: name,
		IconFilePath: iconFilePath,
		PrimaryColor: config.DefaultPrimaryPhaddergruppColor,
		SecondaryColor: config.DefaultSecondaryPhaddergruppColor,
		MumsPrice: 10,
		SwishRecipientNumber: "",
		SwishRecipientName: "",
	}
}

