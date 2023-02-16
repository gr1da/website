package main

import (
	"encoding/json"
	"net/http"
	"os"
)

const apiUrl = `https://api.open-meteo.com/v1/forecast?latitude=51.05&longitude=-114.09&hourly=temperature_2m`

type WeatherResponse struct {
	Hourly struct {
		Time           []string
		Temperature_2m []float64
	}
}

type WeatherProvider interface {
	get() *WeatherResponse
}

type TestWeather struct {
	get func() *WeatherResponse
}

func NewTestWeather() (*TestWeather, error) {
	f, _ := os.Open("test_weather.json")
	_wr := &WeatherResponse{}
	json.NewDecoder(f).Decode(_wr)

	_get := func() *WeatherResponse {
		return _wr
	}

	return &TestWeather{get: _get}, nil
}

type RealWeather struct {
	get func() *WeatherResponse
}

func NewRealWeather() (*RealWeather, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_cont := WeatherResponse{}
	err = json.NewDecoder(resp.Body).Decode(&_cont)
	if err != nil {
		return nil, err
	}

	_get := func() *WeatherResponse {
		return &_cont
	}

	return &RealWeather{get: _get}, nil
}
