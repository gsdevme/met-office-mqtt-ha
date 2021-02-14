package metoffice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type Forecast struct {
	FeelsLikeTemperature string `json:"F"`
	Temperature string `json:"T"`
	Humidity string `json:"H"`
	Weather string `json:"W"`
	PrecipitationProbability string `json:"Pp"`
	WindSpeed string `json:"S"`
	WindGuest string `json:"G"`
	UVIndex string `json:"U"`
	Visibility string `json:"V"`
	WindDirection string `json:"D"`
}

type DayPeriod struct {
	Type string `json:"type"`
	Date string `json:"value"`
	Rep []Forecast `json:"Rep"`
}

type WeatherLocation struct {
	Name string `json:"name"`
	Country string `json:"country"`
	DayPeriod []DayPeriod `json:"Period"`
}

type Config struct {
	ApiKey string
	LocationId int
}

func GetForecast(c Config) (*WeatherLocation, error) {
	var url = fmt.Sprintf("http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/regionalforecast/json/%d?res=3hourly&key=%s", c.LocationId, c.ApiKey)

	body, err := doRequest(url)

	result := gjson.Get(string(body), "SiteRep.DV.Location")
	var w *WeatherLocation

	err = json.Unmarshal([]byte(result.String()), &w)

	return w, err
}

func doRequest(u string) ([]byte, error) {
	resp, err := http.Get(u)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("non-200 response returned")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body, err
}