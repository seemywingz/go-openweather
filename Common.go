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

func GetUnitDistance(unit string) string {
	switch unit {
	case "imperial":
		return "mi"
	default:
		return "m"
	}
}

func GetUnitSpeed(unit string) string {
	switch unit {
	case "imperial":
		return "mph"
	default:
		return "m/s"
	}
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

func GetIcon(icon string) string {
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
		return "🌨" // snow
	case "50d", "50n":
		return "🌫" // fog
	default:
		return ""
	}
}

func GetMoonPhaseIcon(phase float64) string {
	switch {
	case phase == 0:
		return "🌑" // New Moon
	case phase > 0 && phase < 0.25:
		return "🌒" // Waxing Crescent
	case phase == 0.25:
		return "🌓" // First Quarter
	case phase > 0.25 && phase < 0.5:
		return "🌔" // Waxing Gibbous
	case phase == 0.5:
		return "🌕" // Full Moon
	case phase > 0.5 && phase < 0.75:
		return "🌖" // Waning Gibbous
	case phase == 0.75:
		return "🌗" // Last Quarter
	case phase > 0.75 && phase < 1:
		return "🌘" // Waning Crescent
	default:
		return ""
	}
}

func GetTimeZoneTLA(timezone string) string {
	switch timezone {
	case "America/New_York":
		return "EST"
	case "America/Chicago":
		return "CST"
	case "America/Denver":
		return "MST"
	case "America/Los_Angeles":
		return "PST"
	case "America/Anchorage":
		return "AKST"
	case "Pacific/Honolulu":
		return "HST"
	case "America/Phoenix":
		return "MST"
	case "Europe/London":
		return "GMT"
	case "Europe/Berlin":
		return "CET"
	case "Asia/Tokyo":
		return "JST"
	case "Australia/Sydney":
		return "AEST"
	case "Asia/Kolkata":
		return "IST"
	case "Asia/Shanghai":
		return "CST"
	case "America/Sao_Paulo":
		return "BRT"
	case "Africa/Johannesburg":
		return "SAST"
	case "Pacific/Auckland":
		return "NZST"
	default:
		return timezone
	}
}
