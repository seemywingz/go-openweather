package openWeather

import (
	"fmt"
	"os"
)

const apiURL = "https://api.openweathermap.org/data/3.0/onecal"

var OPEN_WEATHER_MAP_API_KEY string

func getAPIKey() error {
	OPEN_WEATHER_MAP_API_KEY = os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	if OPEN_WEATHER_MAP_API_KEY == "" {
		return fmt.Errorf("OPEN_WEATHER_MAP_API_KEY is not set")
	}
	return nil
}

func Now() {
	fmt.Println("wAether 🌤️ in you CLI")
}
