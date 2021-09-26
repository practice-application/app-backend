package model

import (
	"github.com/biter777/countries"
)

type Location struct {
	StreetNumber int                   `json:"streetNumber"`
	StreetName   string                `json:"streetName"`
	Suburb       string                `json:"suburb"`
	City         string                `json:"city"`
	Country      countries.CountryCode `json:"contry"`
}
