package models

// *NOTE: integer values are cuurently 4 bytes, may need to increase

// Root type for weather data type
type Weather struct {
	Lon        float64        `json:"lon"`
	Lat        float64        `json:"lot"`
	Weather    []weatherChild `json:"weather"`
	Base       string         `json:"base"`
	Main       mainChild      `json:"main"`
	Visibility int            `json:"visibility"`
	Wind       windChild      `json:"wind"`
	CityId     int32          `json:"id"`
	City       string         `json:"name"`
}

// Weather child representing weather condition data
type weatherChild struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Main child representing tempurature data
type mainChild struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// Wind child representing wind data
type windChild struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Cloud child respresenting cloud data
type coudChild struct {
	All int `json:"all"`
}
