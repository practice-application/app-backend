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
	ID          string    `json:"id"`
	Author      string    `json:"author"`
	Caption     string    `json:"caption"`
	ContentType string    `json:"contentType"`
	DateTime    time.Time `json:"dateTime"`
	FileID      string    `json:"fileID"`
	FileSize    int64     `json:"fileSize"`
	Height      int       `json:"height"`
	Name        string    `json:"name"`
	Width       int       `json:"width"`
}
