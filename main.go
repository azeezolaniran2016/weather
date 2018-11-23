package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	apiURL        = "https://weather.theonlinekenyan.com/api/forecast/"
	apiOkResponse = 200
)

type data struct {
	Day        string `json:"Day"`
	ForcastDay struct {
		Conditions string `json:"Conditions"`
		High       struct {
			Celsius string `json:"Celsius"`
		} `json:"High"`
		Low struct {
			Celsius string `json:"Celsius"`
		} `json:"Low"`
	} `json:"ForecastDay"`
}

type response struct {
	Status  int     `json:"Status"`
	Success bool    `json:"Success"`
	Message string  `json:"Message"`
	Data    []*data `json:"Data"`
}

func main() {
	town, err := parseTown()
	if err != nil {
		log.Fatal(err.Error())
	}

	endpoint := apiURL + strings.ToLower(town)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Fatalf("failed to create request: %s", err.Error())
	}

	client := http.DefaultClient
	client.Timeout = 10 * time.Second

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed request to %s: %s", endpoint, err.Error())
	}

	reqResponse := &response{}
	err = json.NewDecoder(res.Body).Decode(reqResponse)
	res.Body.Close()

	if err != nil {
		log.Fatalf("failed to decode response body: %s", err.Error())
	}

	if reqResponse.Status != apiOkResponse {
		log.Fatalf("received unexpected response status. expected %d, got %d for %s town", apiOkResponse, reqResponse.Status, town)
	}

	data := reqResponse.Data[0]
	const shortForm = "2006-01-02 00:00:00"
	t, err := time.Parse(shortForm, data.Day)
	if err != nil {
		log.Fatalf("Failed to parse day: %s", err.Error())
	}

	fmt.Printf("Town: %s\n", town)
	fmt.Printf("Date: %s\n", t.Format("02 Jan 2006"))

	fmt.Printf("Conditions: %s\n", data.ForcastDay.Conditions)
	fmt.Printf("Temprature: %s Celsius / %s Celsius\n", data.ForcastDay.High.Celsius, data.ForcastDay.Low.Celsius)
}

func parseTown() (string, error) {
	townArg := os.Args[1]
	townVal := os.Args[2]
	if townArg != "-town" {
		return "", fmt.Errorf("unknown argument passed %s. Argument should be in format `weather -town {Town}`", townArg)
	}
	return townVal, nil
}
