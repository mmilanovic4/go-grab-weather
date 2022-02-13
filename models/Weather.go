package models

import (
	"fmt"
	"time"
)

type WeatherResponse struct {
	Coordinates struct {
		Latitude  float32 `json:"lat"`
		Longitude float32 `json:"lon"`
	} `json:"coord"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temperature    float32 `json:"temp"`
		TemperatureMin float32 `json:"temp_min"`
		TemperatureMax float32 `json:"temp_max"`
		Humidity       int     `json:"humidity"`
		Pressure       int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed   float32 `json:"speed"`
		Degrees float32 `json:"deg"`
	} `json:"wind"`
	System struct {
		CountryCode string `json:"country"`
		Sunrise     int64  `json:"sunrise"`
		Sunset      int64  `json:"sunset"`
	} `json:"sys"`
	DateTime int64   `json:"dt"`
	Timezone float64 `json:"timezone"`
	CityName string  `json:"name"`
}

type Weather struct {
	Location              string
	Sunrise               string
	Sunset                string
	TimeOfDataCalculation string
	TemperatureCurrent    string
	TemperatureMinimum    string
	TemperatureMaximum    string
	Humidity              string
	Pressure              string
}

func (wr WeatherResponse) GetWeather() Weather {
	return Weather{
		Location:              fmt.Sprintf("%s %s [%f, %f]", wr.CityName, wr.System.CountryCode, wr.Coordinates.Latitude, wr.Coordinates.Longitude),
		Sunrise:               FormattedDate(wr.System.Sunrise, false),
		Sunset:                FormattedDate(wr.System.Sunset, false),
		TimeOfDataCalculation: FormattedDate(wr.DateTime, true),
		TemperatureCurrent:    fmt.Sprintf("%.1f °C", wr.Main.Temperature),
		TemperatureMinimum:    fmt.Sprintf("%.1f °C", wr.Main.TemperatureMin),
		TemperatureMaximum:    fmt.Sprintf("%.1f °C", wr.Main.TemperatureMax),
		Humidity:              fmt.Sprintf("%d%%", wr.Main.Humidity),
		Pressure:              fmt.Sprintf("%d hPa", wr.Main.Pressure),
	}
}

func FormattedDate(ts int64, verbose bool) string {
	t := time.Unix(ts, 0)
	short := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	if verbose {
		return fmt.Sprintf("%02d.%02d.%d %s", t.Day(), t.Month(), t.Year(), short)
	}
	return short
}
