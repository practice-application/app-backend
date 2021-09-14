package data

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

type Org struct {
	ID               string   `json:"id"`
	OrganisationName string   `json:"organisationName"`
	OrgType          string   `json:"orgType"`
	OrgSize          string   `json:"orgSize"`
	People           []Person `json:"people"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
