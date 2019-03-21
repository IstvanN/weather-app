package main

import "testing"

var testData = rawData{
	WeatherData{
		Predictability: 1,
	},
	WeatherData{
		Predictability: 2,
	},
	WeatherData{
		Predictability: 3,
	},
}

func TestGetHighestPredict(t *testing.T) {
	if getHighestPredict(testData).Predictability != 3 {
		t.Fatalf("getHighestPredict failed: want %d, got %d", getHighestPredict(testData).Predictability, 3)
	}
}
