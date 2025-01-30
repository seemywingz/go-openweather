package openWeather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var OPEN_WEATHER_MAP_API_KEY string

const apiURL = "https://api.openweathermap.org/data/3.0/onecal"

func Init() error {
	OPEN_WEATHER_MAP_API_KEY = os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	if OPEN_WEATHER_MAP_API_KEY == "" {
		return fmt.Errorf("OPEN_WEATHER_MAP_API_KEY is not set")
	}

	return nil
}

func Now() {
	fmt.Println("wAether 🌤️ in you CLI")
}

// get lat and long from provided town, city or zip code
func GetGeoData(location string) (GeoData, error) {
	url := "http://api.openweathermap.org/geo/1.0/direct?q=" + location + "&limit=1&appid=" + OPEN_WEATHER_MAP_API_KEY

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get geo data: %s", resp.Status)
	}

	var geoData GeoData
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return nil, err
	}

	return geoData, nil
}
