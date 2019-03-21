package main

var wd WeatherData

// WeatherData is the aggregated data structure for the received JSON format of the API
type WeatherData struct {
	CityName   string `json:"title"`
	State      string `json:"weather_state_name"`
	MinTemp    int    `json:"min_temp"`
	MaxTemp    int    `json:"max_temp"`
	ActualTemp int    `json:"the_temp"`
}
