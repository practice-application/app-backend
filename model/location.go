package model

import (
	"github.com/biter777/countries"
)

type Location struct {
	AddressLine1 int                   `json:"addressLine1"`
	AddressLine2 string                `json:"addressLine2"`
	Suburb       string                `json:"suburb"`
	City         string                `json:"city"`
	Country      countries.CountryCode `json:"contry"`
}
