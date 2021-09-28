package model

import (
	"time"

	"github.com/biter777/countries"
)

type Person struct {
	ID           string            `json:"id"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	BirthDate    string            `json:"birthDate"`
	Email        string            `json:"email"`
	Phone        string            `json:"phone"`
	Product      []Product         `json:"product"`
	Date         time.Time         `json:"date"`
	AddressLine1 string            `json:"addressLine1"`
	AddressLine2 string            `json:"addressLine2"`
	Suburb       string            `json:"suburb"`
	City         string            `json:"city"`
	Region       string            `json:"region"`
	Country      countries.Country `json:"country"`
}

type Page struct {
	Data    []Person `json:"data"`
	Matches int64    `json:"matches"`
}
