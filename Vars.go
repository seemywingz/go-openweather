package openWeather

import (
	"fmt"
	"os"
)

const apiURL = "https://api.openweathermap.org/data/3.0/onecall"

var apiKey string

func getAPIKey() error {
	apiKey = os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPEN_WEATHER_MAP_API_KEY is not set")
	}
	return nil
}
