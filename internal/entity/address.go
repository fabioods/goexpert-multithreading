package entity

type Address struct {
	ZipCode       string `json:"zipcode"`
	Street        string `json:"street"`
	Neighbourhood string `json:"neighbourhood"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
}
