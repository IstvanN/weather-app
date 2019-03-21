package main

var wd WeatherData

type rawData struct {
	ConsolidatedWeather []struct {
		State          string  `json:"weather_state_name"`
		MinTemp        float64 `json:"min_temp"`
		MaxTemp        float64 `json:"max_temp"`
		ActualTemp     float64 `json:"the_temp"`
		Predictability int     `json:"predictability"`
	} `json:"consolidated_weather"`
	Title string `json:"title"`
}

// WeatherData is the aggregated data structure for the received JSON format of the API
type WeatherData struct {
	CityName       string  `json:"title"`
	State          string  `json:"weather_state_name"`
	MinTemp        float64 `json:"min_temp"`
	MaxTemp        float64 `json:"max_temp"`
	ActualTemp     float64 `json:"the_temp"`
	Predictability int     `json:"predictability"`
}

func getHighestPredict(raw rawData) WeatherData {
	var valueOfHighestPredict int
	var indexOfHighestPredict int

	for i, wd := range raw.ConsolidatedWeather {
		if wd.Predictability > valueOfHighestPredict {
			valueOfHighestPredict = wd.Predictability
			indexOfHighestPredict = i
		}
	}

	highestPredictData := raw.ConsolidatedWeather[indexOfHighestPredict]

	highestWeatherData := WeatherData{
		CityName:       raw.Title,
		State:          highestPredictData.State,
		MinTemp:        highestPredictData.MinTemp,
		MaxTemp:        highestPredictData.MaxTemp,
		ActualTemp:     highestPredictData.ActualTemp,
		Predictability: highestPredictData.Predictability,
	}
	return highestWeatherData
}
