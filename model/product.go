package model

import (
	"time"
)

type Product struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Price       string        `json:"price"`
	Description string        `json:"description"`
	Date        time.Time     `json:"date"`
	ImageID     string        `json:"imageID"`
	User        string        `json:"user"`
	Tags        [0 & 5]string `json:"tags"`
	Category    string        `json:"category"`
}

type ProductPage struct {
	Data    []Product `json:"data"`
	Matches int64     `json:"matches"`
}
