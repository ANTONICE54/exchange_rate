package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type SecRateProvider struct {
}

func NewSecRateProvider() *SecRateProvider {
	return &SecRateProvider{}
}

type SecRateProviderResponse struct {
	Rate string `json:"rate"`
}

func (rateP *SecRateProvider) GetRate() (*float64, error) {
	apiKey := "46af2b81f2"
	url := fmt.Sprintf("https://www.live-rates.com/api/price?key=%s&rate=USD_UAH", apiKey)

	resp, err := http.Get(url)
	if err != nil {

		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	var fetchedObject []SecRateProviderResponse

	err = json.Unmarshal(body, &fetchedObject)
	if err != nil {

		return nil, err
	}

	rate, err := strconv.ParseFloat(fetchedObject[0].Rate, 64)
	if err != nil {

		return nil, err
	}
	log.Printf("live-rates.com - got value from provider %v", rate)

	return &rate, nil
}
