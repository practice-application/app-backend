package model

import (
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Image       Image     `json:"image"`
}

type ProductPage struct {
	Data    []Product `json:"data"`
	Matches int64     `json:"matches"`
}

type Image struct {
	DataUrl          string    `json:"data_url"`
	Type             string    `json:"type"`
	LastModified     int64     `json:"lastModified"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
	Size             int64     `json:"fileSize"`
	Name             string    `json:"name"`
}
