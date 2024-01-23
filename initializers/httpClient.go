package initializers

import (
	"net/http"
)

var HttpClient http.Client

func CreateHttpClient() {
	HttpClient = http.Client{}
}
