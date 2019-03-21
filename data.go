package main

var wd WeatherData

type rawData []WeatherData

// WeatherData is the aggregated data structure for the received JSON format of the API
type WeatherData struct {
	CityName       string `json:"title"`
	State          string `json:"weather_state_name"`
	MinTemp        int    `json:"min_temp"`
	MaxTemp        int    `json:"max_temp"`
	ActualTemp     int    `json:"the_temp"`
	Predictability int    `json:"predictability"`
}

func getHighestPredict(raw rawData) WeatherData {
	var valueOfHighestPredict int
	var indexOfHighestPredict int

	for i, wd := range raw {
		if wd.Predictability > valueOfHighestPredict {
			valueOfHighestPredict = wd.Predictability
			indexOfHighestPredict = i
		}
	}

	highestPredictData := raw[indexOfHighestPredict]

	highestWeatherData := WeatherData{
		CityName:       highestPredictData.CityName,
		State:          highestPredictData.State,
		MinTemp:        highestPredictData.MinTemp,
		MaxTemp:        highestPredictData.MaxTemp,
		ActualTemp:     highestPredictData.ActualTemp,
		Predictability: highestPredictData.Predictability,
	}
	return highestWeatherData
}
