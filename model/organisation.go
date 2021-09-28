package model

import "github.com/biter777/countries"

type Org struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	OrgType      string            `json:"orgType"`
	OrgSize      string            `json:"orgSize"`
	People       []Person          `json:"people"`
	AddressLine1 string            `json:"addressLine1"`
	AddressLine2 string            `json:"addressLine2"`
	Suburb       string            `json:"suburb"`
	City         string            `json:"city"`
	Region       string            `json:"region"`
	Country      countries.Country `json:"country"`
}
