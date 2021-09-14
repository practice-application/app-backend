package model

type Person struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       string    `json:"age"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Product   []Product `json:"product"`
	Date      string    `json:"date"`
}

