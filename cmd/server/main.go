package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fabioods/goexpert-multithreading/internal/dto"
	"github.com/fabioods/goexpert-multithreading/internal/entity"
)

const timeOut = 1 * time.Second

func fetchFromBrasilApi(ctx context.Context, cep string, channelResponse chan string, channelAddres chan entity.Address) {
	if cep == "" {
		channelResponse <- "No cep provided"
		return
	}

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		channelResponse <- "Failed to get request by Brasil API"
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		channelResponse <- "Failed to get response by Brasil API"
		return
	}

	defer resp.Body.Close()

	var address dto.AddressBrasilApi
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		channelResponse <- "Failed to decode response by Brasil API"
		return
	}

	channelResponse <- "Brasil API"
	channelAddres <- address.ToAddress()

}

func fetchFromViaCep(ctx context.Context, cep string, channelResponse chan string, channelAddres chan entity.Address) {
	if cep == "" {
		channelResponse <- "No cep provided"
		return
	}

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		channelResponse <- "Failed to get request by Via Cep"
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		channelResponse <- "Failed to get response by Via Cep"
		return
	}

	defer resp.Body.Close()

	var address dto.AddressViaCep
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		channelResponse <- "Failed to decode response by Via Cep"
		return
	}

	channelResponse <- "Via Cep"
	channelAddres <- address.ToAddress()

}

func main() {
	cep := "84030270"

	channelResponse := make(chan string)
	channelAddress := make(chan entity.Address)

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	go fetchFromViaCep(ctx, cep, channelResponse, channelAddress)
	go fetchFromBrasilApi(ctx, cep, channelResponse, channelAddress)

	select {
	case chanResp := <-channelResponse:
		address := <-channelAddress
		fmt.Println("Responde from: ", chanResp)
		adressFormated := fmt.Sprintf("CEP: %s, Rua: %s, Bairro: %s, Cidade: %s, Estado: %s", address.ZipCode, address.Street, address.Neighbourhood, address.City, address.State)
		fmt.Println(adressFormated)
	case <-ctx.Done():
		fmt.Println("Timeout")
	}

}
