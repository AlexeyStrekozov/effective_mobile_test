package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/AlexeyStrekozov/effective_mobile_test/initializers"
)

const nationalizeUrl = "https://api.nationalize.io/?name="

type nationalityData struct {
	Country []countryData
}

type countryData struct {
	Name string `json:"country_id"`
}

func GetNationality(name string) string {
	resp, err := initializers.HttpClient.Get(nationalizeUrl + name)

	if err != nil {
		fmt.Println("Nationalize is't working ")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Nationalize: Failed to parse")
	}

	var data nationalityData

	json.Unmarshal(body, &data)

	return data.Country[0].Name
}
