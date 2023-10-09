package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ViaCepSt struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilCepSt struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Location     struct {
		Type        string `json:"type"`
		Coordinates struct {
			Longitude string `json:"longitude"`
			Latitude  string `json:"latitude"`
		} `json:"coordinates"`
	} `json:"location"`
}

type BrasilAbertoSt struct {
	Meta struct {
		CurrentPage  int `json:"currentPage"`
		ItemsPerPage int `json:"itemsPerPage"`
		TotalOfItems int `json:"totalOfItems"`
		TotalOfPages int `json:"totalOfPages"`
	} `json:"meta"`
	Result struct {
		Street         string `json:"street"`
		Complement     string `json:"complement"`
		District       string `json:"district"`
		DistrictID     int    `json:"districtId"`
		City           string `json:"city"`
		CityID         int    `json:"cityId"`
		IbgeID         int    `json:"ibgeId"`
		State          string `json:"state"`
		StateShortname string `json:"stateShortname"`
		Zipcode        string `json:"zipcode"`
	} `json:"result"`
}

func ApiCep[T any](url string, apiStruct T, result chan T) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error in request: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
	}
	err = json.Unmarshal(body, &apiStruct)
	if err != nil {
		log.Printf("Error parsing json: %v", err)
	}
	result <- apiStruct
}

func main() {
	api1 := make(chan ViaCepSt)
	api2 := make(chan BrasilCepSt)
	api3 := make(chan BrasilAbertoSt)

	var viaCep ViaCepSt
	var BrasilCep BrasilCepSt
	var BrasilAberto BrasilAbertoSt
	go ApiCep("https://viacep.com.br/ws/01153000/json/", viaCep, api1)
	go ApiCep("https://brasilapi.com.br/api/cep/v2/01153000", BrasilCep, api2)
	go ApiCep("https://brasilaberto.com/api/v1/zipcode/01153000", BrasilAberto, api3)

	select {
	case result1 := <-api1:
		log.Fatalf("Received from ViaCep: \n %v\n", result1)
	case result2 := <-api2:
		log.Fatalf("Received from BrasilCep: \n %v\n", result2)
	case result3 := <-api3:
		log.Fatalf("Received from BrasilAberto: \n %v\n", result3)
	case <-time.After(time.Second * 2):
		log.Fatalf("Timeout")
	}
}
