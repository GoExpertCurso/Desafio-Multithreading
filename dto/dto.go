package dto

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
