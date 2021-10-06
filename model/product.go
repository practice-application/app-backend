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
	ID      string `json:"id"`
	DataUrl string `json:"data_url"`
	Name    string `json:"name"`
}
