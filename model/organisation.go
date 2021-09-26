package model

type Org struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	OrgType string   `json:"orgType"`
	OrgSize string   `json:"orgSize"`
	People  []Person `json:"people"`
	Address Location `json:"address"`
}
