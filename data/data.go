package data

type Person struct {
	ID        string    `bson:"id"`
	FirstName string    `bson:"firstName"`
	LastName  string    `bson:"lastName"`
	Age       int       `bson:"age"`
	Email     string    `bson:"email"`
	Phone     string    `bson:"phone"`
	Product   []Product `bson:"product"`
}

type Org struct {
	ID               string   `bson:"id"`
	OrganisationName string   `bson:"organisationName"`
	OrgType          string   `bson:"orgType"`
	OrgSize          string   `bson:"orgSize"`
	People           []Person `bson:"people"`
}

type Product struct {
	ID          string `bson:"id"`
	Name        string `bson:"name"`
	Price       int    `bson:"price"`
	Description string `bson:"description"`
}
