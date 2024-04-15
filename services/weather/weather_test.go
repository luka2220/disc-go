package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/luka2220/discGo/models"
	"github.com/stretchr/testify/assert"
)

func testWeatherHandler(w http.ResponseWriter, req *http.Request) {
	weatherTest := &models.Weather{
		City:   "Mississauga",
		CityId: 10110,
		Lon:    43.5890,
		Lat:    79.6441,
		Weather: []models.WeatherChild{
			{
				Id:   23,
				Main: "Sunny",
			},
		},
		Main: models.MainChild{
			Temp: 19.21,
		},
	}

	byteResponse, _ := json.Marshal(weatherTest)
	w.Write(byteResponse)
}

func createHTTPServer() {
	http.HandleFunc("/", testWeatherHandler)

	if err := http.ListenAndServe(":42069", nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

func TestWeatherService(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go createHTTPServer()

	testURL := "http://127.0.0.1:42069"

	testOpt := &WeatherDataService{
		testURL,
	}

	weatherService := NewWeatherService("", *testOpt)
	weatherData, _ := weatherService.FetchWeatherData()

	log.Println("Weather data (client) = ", *weatherData)

	// *TODO: Asserts
	weatherAssert := &models.Weather{
		City:   "Mississauga",
		CityId: 10110,
		Lon:    43.5890,
		Lat:    79.6441,
		Weather: []models.WeatherChild{
			{
				Id:   23,
				Main: "Sunny",
			},
		},
		Main: models.MainChild{
			Temp: 19.21,
		},
	}

	assert.Equal(t, *weatherAssert, *weatherData, "Should be equal...")

}
