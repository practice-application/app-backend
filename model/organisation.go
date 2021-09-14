package model

type Org struct {
	ID               string   `json:"id"`
	OrganisationName string   `json:"organisationName"`
	OrgType          string   `json:"orgType"`
	OrgSize          string   `json:"orgSize"`
	People           []Person `json:"people"`
}

