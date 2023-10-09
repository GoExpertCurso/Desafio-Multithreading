package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	dto "github.com/lucianosz7/GoExpert/MULTITHREADING/dto"
)

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
	api1 := make(chan dto.ViaCepSt)
	api2 := make(chan dto.BrasilCepSt)
	api3 := make(chan dto.BrasilAbertoSt)

	var viaCep dto.ViaCepSt
	var BrasilCep dto.BrasilCepSt
	var BrasilAberto dto.BrasilAbertoSt
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
