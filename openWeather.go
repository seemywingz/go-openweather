package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OneCallResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Clouds     int     `json:"clouds"`
		Uvi        float64 `json:"uvi"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindGust   float64 `json:"wind_gust"`
		WindDeg    int     `json:"wind_deg"`
		Rain       struct {
			OneH float64 `json:"1h"`
		} `json:"rain"`
		Snow struct {
			OneH float64 `json:"1h"`
		} `json:"snow"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
	Minutely []struct {
		Dt            int     `json:"dt"`
		Precipitation float64 `json:"precipitation"`
	} `json:"minutely"`
	Hourly []struct {
		Dt         int     `json:"dt"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindGust   float64 `json:"wind_gust"`
		WindDeg    int     `json:"wind_deg"`
		Pop        float64 `json:"pop"`
		Rain       struct {
			OneH float64 `json:"1h"`
		} `json:"rain"`
		Snow struct {
			OneH float64 `json:"1h"`
		} `json:"snow"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"hourly"`
	Daily []struct {
		Dt        int     `json:"dt"`
		Sunrise   int     `json:"sunrise"`
		Sunset    int     `json:"sunset"`
		Moonrise  int     `json:"moonrise"`
		Moonset   int     `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Summary   string  `json:"summary"`
		Temp      struct {
			Morn  float64 `json:"morn"`
			Day   float64 `json:"day"`
			Eve   float64 `json:"eve"`
			Night float64 `json:"night"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
		} `json:"temp"`
		FeelsLike struct {
			Morn  float64 `json:"morn"`
			Day   float64 `json:"day"`
			Eve   float64 `json:"eve"`
			Night float64 `json:"night"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindGust  float64 `json:"wind_gust"`
		WindDeg   int     `json:"wind_deg"`
		Clouds    int     `json:"clouds"`
		Uvi       float64 `json:"uvi"`
		Pop       float64 `json:"pop"`
		Rain      float64 `json:"rain"`
		Snow      float64 `json:"snow"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"daily"`
	Alerts []struct {
		SenderName  string   `json:"sender_name"`
		Event       string   `json:"event"`
		Start       int      `json:"start"`
		End         int      `json:"end"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	} `json:"alerts"`
}

func Get(lat, long float64, unit, apiKey string) (OneCallResponse, error) {

	url := fmt.Sprintf("%s?lat=%f&lon=%f&exclude=minutely&units=%s&appid=%s", apiURL, lat, long, unit, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return OneCallResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OneCallResponse{}, fmt.Errorf("error reading response body: %w", err)
	}
	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return OneCallResponse{}, fmt.Errorf("unexpected status: %s, body: %s", resp.Status, string(body))
	}
	// Decode the JSON data
	var weatherData OneCallResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return OneCallResponse{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	return weatherData, nil
}
