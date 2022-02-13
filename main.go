package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mmilanovic4/go-grab-weather/models"
	"github.com/mmilanovic4/go-grab-weather/utils"
)

const URL = "https://api.openweathermap.org/data/2.5/weather"

var (
	query   *string
	verbose *bool
)

func main() {
	query = flag.String("q", "", "City")
	verbose = flag.Bool("v", false, "Verbosity")
	flag.Parse()

	params := make(map[string]string)
	params["q"] = *query
	params["units"] = "metric"
	params["appid"] = os.Getenv("OPEN_WEATHER_API_KEY")
	output := &models.WeatherResponse{}

	utils.SendRequest(URL, params, output, *verbose)

	weather := output.GetWeather()

	fmt.Printf("%s\n%s\n\n", weather.Location, weather.TimeOfDataCalculation)
	fmt.Printf("%-30s %s\n", "Current temperature:", weather.TemperatureCurrent)
	fmt.Printf("%-30s %s\n", "Minimum temperature:", weather.TemperatureMinimum)
	fmt.Printf("%-30s %s\n\n", "Maximum temperature:", weather.TemperatureMaximum)
	fmt.Printf("%-30s %s\n", "Humidity:", weather.Humidity)
	fmt.Printf("%-30s %s\n\n", "Pressure:", weather.Pressure)
	fmt.Printf("%-30s %s\n", "Sunrise:", weather.Sunrise)
	fmt.Printf("%-30s %s\n", "Sunset:", weather.Sunset)
}
