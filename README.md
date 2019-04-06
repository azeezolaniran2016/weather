# weather

weather is a simple cli application which allows users query weather data for a specific location. It was written as part of a live coding session in a tech assessment.

## Intro

An API endpoint exists at `https://weather.theonlinekenyan.com/api/` that will give weather forecasts for most kenyan towns.
API request format: `https://weather.theonlinekenyan.com/api/forecast/TOWN` eg
for Nairobi: `https://weather.theonlinekenyan.com/api/forecast/nairobi` for Lodwar: `https://weather.theonlinekenyan.com/api/forecast/lodwar` for Rongai: `https://weather.theonlinekenyan.com/api/forecast/rongai`
Note: No API request authentication is required
The API responds with a JSON payload in the following format:

```
{
  "Status": 200,
  "Success": true,
  "Message": "OK",
  "Data": [{
    "Day": "2018-11-22 00:00:00",
    "ForecastDay": {
      "Date": {
        "Epoch": "1542902400",
        "Day": 22,
        "Month": 11,
        "Year": 2018,
        "weekday_short": "Thu",
        "tz_short": "EAT"
      },
      "Period": 1,
      "High": {
        "Celsius": "37"
      },
      "Low": {
        "Celsius": "25"
      },
      "Conditions": "Partly Cloudy",
      "Icon": "partlycloudy",
      "SkyIcon": "",
      "Pop": 0,
      "MaxWind": {
        "KPH": 40,
        "Dir": "E",
        "Degrees": 93
      },
      "AveWind": {
        "KPH": 27,
        "Dir": "E",
        "Degrees": 93
      },
      "AveHumidity": 22
    }
  }]
}
```
        
### Task

Write a command line application called weather, that accepts an argument specifying a town. It should then print the town’s weather Conditions and Temperature (High/Low)
eg. weather -town Rongai should return
```
Town: Rongai
Date: 22 Nov 2018
Conditions: Chance of a Thunderstorm
Temperature: 25 Celsius / 13 Celsius
```

### How to run

You can run the program using the `go run` command or compile and run the resulting executable:

- With `go run` for city named Rongai:
```
go run main.go -city Rongai
```

- Compile exectuable name weather, and run it for city name Rongai:

```
go compile -o weather && weather -city Rongai
```