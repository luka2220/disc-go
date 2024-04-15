package weather

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/luka2220/discGo/models"
)

type WeatherDataService struct {
	URL string
}

func NewWeatherService(city string, params ...WeatherDataService) *WeatherDataService {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=4b131da2c3d34ffbf959576862e58b66&units=metric", city)

	if len(params) > 0 {
		url = params[0].URL
	}

	return &WeatherDataService{
		url,
	}
}

func (w *WeatherDataService) FetchWeatherData() (*models.Weather, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	response, err := http.Get(w.URL)
	if err != nil {
		log.Panic("An error occured fetching weather data", err)
		return nil, err
	}

	// Close the body once the function execution returns
	defer response.Body.Close()

	log.Printf("Response status: %s\n", response.Status)

	scanner := bufio.NewScanner(response.Body)

	var weatherDataBytes []byte
	var weatherData *models.Weather

	for i := 0; scanner.Scan() && i < 5; i++ {
		weatherDataBytes = scanner.Bytes()
	}

	if err := json.Unmarshal(weatherDataBytes, &weatherData); err != nil {
		log.Panic("Error occured when unmarshalling data", err)
		return nil, err
	}

	return weatherData, nil
}
