package model

import (
	"time"
)

type Person struct {
	ID           string    `json:"id"`
	Auth0ID      string    `json:"auth0id"`
	UserName     string    `json:"userName"`
	Avatar       string    `json:"avatar"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	BirthDate    string    `json:"birthDate"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Product      []Product `json:"product"`
	Date         time.Time `json:"date"`
	AddressLine1 string    `json:"addressLine1"`
	AddressLine2 string    `json:"addressLine2"`
	Suburb       string    `json:"suburb"`
	City         string    `json:"city"`
	Region       string    `json:"region"`
	Country      string    `json:"country"`
	Verified     bool      `json:"verified"`
}

type Page struct {
	Data    []Person `json:"data"`
	Matches int64    `json:"matches"`
}
