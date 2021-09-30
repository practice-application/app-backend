package model

import (
	"image"
	"time"
)

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Price       string      `json:"price"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
	Image       image.Alpha `json:"image"`
}

type ProductPage struct {
	Data    []Product `json:"data"`
	Matches int64     `json:"matches"`
}
