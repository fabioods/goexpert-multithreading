package dto

import (
	"github.com/fabioods/goexpert-multithreading/internal/entity"
)

type AddressBrasilApi struct {
	Cep     string `json:"cep"`
	Estado  string `json:"state"`
	Cidade  string `json:"city"`
	Bairro  string `json:"neighborhood"`
	Rua     string `json:"street"`
	Serviço string `json:"service"`
}

func (a AddressBrasilApi) ToAddress() entity.Address {
	return entity.Address{
		ZipCode:       a.Cep,
		Street:        a.Rua,
		Neighbourhood: a.Bairro,
		City:          a.Cidade,
		State:         a.Estado,
	}
}
