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
	ImageID     string    `json:"imageID"`
	User        string    `json:"user"`
	UserPic     string    `json:"userPic"`
	Auth0id     string    `json:"auth0id"`
	Tags        [5]string `json:"tags"`
	Category    string    `json:"category"`
	Size        string    `json:"size"`
	Priority    bool      `json:"priority"`
	Edits       int       `json:"edits"`
}

type ProductPage struct {
	Data    []Product `json:"data"`
	Matches int64     `json:"matches"`
}
