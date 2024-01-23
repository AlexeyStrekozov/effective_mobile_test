package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/AlexeyStrekozov/effective_mobile_test/initializers"
)

const genderizeUrl = "https://api.genderize.io/?name="

type genderizeData struct {
	Count  int
	Name   string
	Gender string
}

func GetGender(name string) string {
	resp, err := initializers.HttpClient.Get(genderizeUrl + name)

	if err != nil {
		fmt.Println("Genderize is't working ")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Genderize: Failed to parse")
	}

	var data genderizeData

	json.Unmarshal(body, &data)

	return data.Gender
}
