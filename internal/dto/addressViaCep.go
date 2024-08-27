package dto

import "github.com/fabioods/goexpert-multithreading/internal/entity"

type AddressViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:" siafi"`
}

func (a AddressViaCep) ToAddress() entity.Address {
	return entity.Address{
		ZipCode:       a.Cep,
		Street:        a.Logradouro,
		Neighbourhood: a.Bairro,
		City:          a.Localidade,
		State:         a.Uf,
	}
}
