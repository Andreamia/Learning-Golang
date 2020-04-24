// Don't forget to do before run the program:
// 	- install "github.com/faiface/beep" package			: 	go get -u github.com/faiface/beep
// 	- install "github.com/faiface/beep/wav3" package	: 	go get -u github.com/faiface/beep/wav
// 	- install "github.com/faiface/beep/speaker" package	: 	go get -u github.com/faiface/beep/speaker
//
// This program:
// 	- get current weather parameters (temperature, pressure, humiduty) for entered city
// from openweathermap.org
// 	- notify user by message and sound alarm if these weather parameters have changed in comparison
// with previous ones
//
// Please, take into account that weather data in "openweathermap.org" is updated
// no more than one time every 10 minutes.
// So to see changes in weather parameters it's recommended to making http request to the API
// no more than one time every 10 minutes for one location

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
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

	// notify if city with cityName wasn't found
	if currentWeather.CityID == 0 && currentWeather.CityName == "" {
		fmt.Println("-----")
		fmt.Printf("There is NO CITY with name = %s!\n", cityName)
		os.Exit(3)
	}

	return currentWeather, err
}

// compare current weather parametrs(temperature, pressure, humiduty) from weatherCurrent
// with previous ones from weatherPrevious
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

// show current weather parameters (temperature, pressure, humiduty)
// notify user if parameter's values have changed: message + sound alarm
func showCurrentWeather(cityName string, weather CurrentWeather, isTemperatureChanged, isPressureChanged, isHumidityChanged bool) {

	fmt.Println("-----")
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
	if isTemperatureChanged == true || isPressureChanged == true || isHumidityChanged == true {
		playAlarm()
	}
}

// delay before next weather request
func delayBeforeNextRequest() error {

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

// download file from url to "alarm.wav" file in current directory
func downloadWAVFileWithAlarm() error {

	var url string = "http://stuffnewspaper.com/sounds/MISC/bell.wav"
	var filePath string

	filePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath += "\\alarm.wav"

	// get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// create the file
	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// play "alarm.wav"
func playAlarm() {

	// downaload file with alarm sound
	err := downloadWAVFileWithAlarm()
	if err != nil {
		panic(err)
	}

	// open file, so that we can decode and play it
	file, err := os.Open("alarm.wav")
	if err != nil {
		panic(err)
	}

	// decode file with wav.Decode to streamer
	streamer, format, err := wav.Decode(file)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	// initialize the speaker
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// play the streamer
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

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

		fmt.Print("\a")
		fmt.Println("Enter 1 or 2 depending on your choice:")
		fmt.Printf("1. Monitor weather in %s\n", city)
		fmt.Printf("2. Exit the program\n")
		fmt.Scanln(&response)
		response = strings.TrimSpace(response)
		fmt.Print("\a")

		switch response {
		// monitor weather in city
		case "1":
			{
				if monitorCounter != 0 {
					err = delayBeforeNextRequest()
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
				showCurrentWeather(city, weather1, isTempChanged, isPresChanged, isHumChanged)
				monitorCounter++
			}
		// exit program
		case "2":
			{
				break MonitorWeather
			}
		default:
			{
				fmt.Println("Incorrect input!")
				break MonitorWeather
			}
		}
	}

}
