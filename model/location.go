package model

import (
	"github.com/biter777/countries"
)

type Location struct {
	AddressLine1 string            `json:"addressLine1"`
	AddressLine2 string            `json:"addressLine2"`
	Suburb       string            `json:"suburb"`
	City         string            `json:"city"`
	Region       string            `json:"region"`
	Country      countries.Country `json:"country"`
}
