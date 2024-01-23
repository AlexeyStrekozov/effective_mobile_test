package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/AlexeyStrekozov/effective_mobile_test/initializers"
)

const agifyUrl = "https://api.agify.io/?name="

type agifyData struct {
	Count int
	Name  string
	Age   int
}

func GetAge(name string) int {
	resp, err := initializers.HttpClient.Get(agifyUrl + name)

	if err != nil {
		fmt.Println("Agify is't working ")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Agify: Failed to parse")
	}

	var data agifyData

	json.Unmarshal(body, &data)

	return data.Age
}
