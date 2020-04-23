package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// constants for getting access to openweather API
const (
	apiURL string = "api.openweathermap.org"
	apiKey string = "358b27409617d93ce1fba388b52aa46f"
)

// Weather from http request json
type Weather struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
}

// CurrentWeather from http request json
type CurrentWeather struct {
	Weather  Weather `json:"main"`
	CityID   int     `json:"id"`
	CityName string  `json:"name"`
}

// get current weather parameters for city wih cityName from http request json
func currentWeather(cityName string) (CurrentWeather, error) {

	var currentWeather CurrentWeather

	// add http client
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	// compile http request and get http response
	url := "http://" + apiURL + "/data/2.5/weather?q=" + cityName + "&units=metric&appid=" + apiKey

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read http response body into a byte stream
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal the byte stream into currentWeather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		panic(err)
	}

	return currentWeather, err
}

// compare current weather parametrs(temperature, pressure, humiduty) with previous ones
func weatherChanged(weatherPrevious, weatherCurrent CurrentWeather) (bool, bool, bool) {

	var isTemperatureChanged bool = false
	var isPressureChanged bool = false
	var isHumidityChanged bool = false

	if math.Abs(weatherCurrent.Weather.Temp-weatherPrevious.Weather.Temp) > 1e-3 {
		isTemperatureChanged = true
	}
	if math.Abs(weatherCurrent.Weather.Pressure-weatherPrevious.Weather.Pressure) > 1e-3 {
		isPressureChanged = true
	}
	if weatherCurrent.Weather.Humidity != weatherPrevious.Weather.Humidity {
		isHumidityChanged = true
	}
	return isTemperatureChanged, isPressureChanged, isHumidityChanged
}

// print current weather parameters (temperature, pressure, humiduty)
// notify user if parameter's values have changed: message + sound
func printCurrentWeather(cityName string, weather CurrentWeather, isTemperatureChanged, isPressureChanged, isHumidityChanged bool) {

	fmt.Println("-----")
	if weather.CityID == 0 && weather.CityName == "" {
		fmt.Printf("There is NO CITY with name = %s!\n", cityName)
		os.Exit(3)
	}
	fmt.Printf("Current weather in %s:\n", cityName)

	if isTemperatureChanged == true {
		fmt.Printf("  Temperature's changed. Temperature = %4.2f °C\n", weather.Weather.Temp)
	} else {
		fmt.Printf("  Temperature = %4.2f °C\n", weather.Weather.Temp)
	}

	if isPressureChanged == true {
		fmt.Printf("  Pressure's changed. Pressure = %6.2f hPa\n", weather.Weather.Pressure)
	} else {
		fmt.Printf("  Pressure = %6.2f hPa\n", weather.Weather.Pressure)
	}

	if isHumidityChanged == true {
		fmt.Printf("  Humidity's changed. Humidity = %3d %%\n", weather.Weather.Humidity)
	} else {
		fmt.Printf("  Humidity = %3d %%\n", weather.Weather.Humidity)
	}
}

// delay before next weather request
func askForDelay() error {

	var delay int
	var strDelay string

	fmt.Println("How long do you want to delay before sending next weather request (enter number of minutes for delay):")
	fmt.Scanln(&strDelay)
	delay, err := strconv.Atoi(strDelay)
	if err != nil {
		fmt.Println("Incorrect input!")
		panic(err)
	}
	fmt.Printf("Next weather request will be sent in %d minutes. Please, wait..\n", delay)
	time.Sleep(time.Duration(delay) * time.Minute)
	return err
}

// program execution
func main() {
	var weather1, weather2 CurrentWeather
	var city string
	var isTempChanged, isPresChanged, isHumChanged bool = false, false, false
	var response string = ""
	var err error
	var monitorCounter int = 0

	// read city name
	fmt.Println("Enter city name:")
	reader := bufio.NewReader(os.Stdin)
	city, _ = reader.ReadString('\n')
	city = strings.TrimSpace(city)

MonitorWeather:
	for true {
		// user decides if he wants to monitor weather in city
		fmt.Println("")
		fmt.Println("Enter 1 or 2 depending on your choice:")
		fmt.Printf("1. Monitor weather in %s\n", city)
		fmt.Printf("2. Exit the program\n")
		fmt.Scanln(&response)
		response = strings.TrimSpace(response)

		switch response {
		// monitor weather in city
		case "1":
			{
				if monitorCounter != 0 {
					err = askForDelay()
					if err != nil {
						panic(err)
					}
				}

				weather2 = weather1
				weather1, err = currentWeather(city)
				if err != nil {
					panic(err)
				}
				if monitorCounter != 0 {
					isTempChanged, isPresChanged, isHumChanged = weatherChanged(weather2, weather1)
				}
				printCurrentWeather(city, weather1, isTempChanged, isPresChanged, isHumChanged)
				monitorCounter++
			}
		// exit program
		case "2":
			{
				break MonitorWeather
			}
		// make correct choice
		default:
			{
				fmt.Println("Incorrect input!")
				break MonitorWeather
			}
		}
	}

}
