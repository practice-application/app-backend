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
}

type ProductPage struct {
	Data    []Product `json:"data"`
	Matches int64     `json:"matches"`
}
