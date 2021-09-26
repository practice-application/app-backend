package model

import (
	"github.com/biter777/countries"
)

type Location struct {
	StreetNumber int                   `json:"stNumber"`
	StreetName   string                `json:"stName"`
	Suburb       string                `json:"suburb"`
	City         string                `json:"city"`
	Country      countries.CountryCode `json:"contry"`
}
