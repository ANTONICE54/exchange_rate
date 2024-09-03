package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RateProvider struct {
}

func NewRateProvider() *RateProvider {
	return &RateProvider{}
}

type RateResponse struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func (rateP *RateProvider) GetRate() (*float64, error) {
	apiKey := "f1dd4d7a12b8ed03793a5728"
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/USD", apiKey)

	resp, err := http.Get(url)
	if err != nil {

		return nil, err

	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rateList RateResponse

	err = json.Unmarshal(body, &rateList)
	if err != nil {
		return nil, err
	}

	rate := rateList.ConversionRates["UAH"]

	log.Printf("exchangerate-api.com - got value from provider %v", rate)

	return &rate, nil
}
