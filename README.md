# OpenWeather client

A simple [OpenWeather API](https://openweathermap.org/) client built in Go for learning purpose.

**Download and build:**

```
git clone https://github.com/mmilanovic4/go-grab-weather.git
cd go-grab-weather
go build ./main.go
```

**Set up your `.envrc`:**

First you'll need to generate your API key on [OpenWeather](https://home.openweathermap.org/api_keys). After that, you'll need to store it in your `.envrc` file:

```
echo export OPEN_WEATHER_API_KEY=1234567890 > my-project/.envrc
direnv allow
echo $OPEN_WEATHER_API_KEY
```

**Run:**

```
clear && ./main -q Belgrade
clear && ./main -q "Los Angeles"
```
