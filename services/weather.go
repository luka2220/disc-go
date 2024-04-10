package services

import (
	"bufio"
	"fmt"
	"github.com/luka2220/discGo/utils"
	"log"
	"net/http"
)

func GetWeatherData() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	city := utils.CityWeather
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=4b131da2c3d34ffbf959576862e58b66&units=metric", city)

	res, err := http.Get(url)
	if err != nil {
		log.Panic("Error occured fetching data", err)
	}

	defer res.Body.Close()

	log.Printf("Response status = %s\n", res.Status)

	scanner := bufio.NewScanner(res.Body)

	log.Println("Scanner text = ", scanner.Text())

	for i := 0; scanner.Scan() && i < 5; i++ {
		log.Println("Scanner text = ", scanner.Text())
	}

}
