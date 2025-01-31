package openWeather

import (
	"fmt"
	"os"
)

const apiURL = "https://api.openweathermap.org/data/3.0/onecall"

var apiKey string

var ValidUnits = map[string]bool{
	"standard": true,
	"metric":   true,
	"imperial": true,
}

func getAPIKey() error {
	apiKey = os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPEN_WEATHER_MAP_API_KEY is not set")
	}
	return nil
}

func GetUnitSymbol(unit string) string {
	switch unit {
	case "standard":
		return "K"
	case "imperial":
		return "°F"
	default:
		return "°C"
	}
}

func GetUnitName(unit string) string {
	switch unit {
	case "standard":
		return "Kelvin"
	case "imperial":
		return "Fahrenheit"
	default:
		return "Celsius"
	}
}

func GetIconEmoji(icon string) string {
	switch icon {
	case "01d":
		return "☀️" // clear-day
	case "01n":
		return "🌙" // clear-night
	case "02d":
		return "🌤" // partly-cloudy-day
	case "02n":
		return "☁🌙" // partly-cloudy-night
	case "03d", "03n":
		return "☁" // cloudy
	case "04d", "04n":
		return "☁" // cloudy
	case "09d", "09n":
		return "🌧" // rain
	case "10d", "10n":
		return "🌦" // rain
	case "11d", "11n":
		return "⛈️" // thunderstorm
	case "13d", "13n":
		return "🌨☃️" // snow
	case "50d", "50n":
		return "🌫" // fog
	default:
		return ""
	}
}
