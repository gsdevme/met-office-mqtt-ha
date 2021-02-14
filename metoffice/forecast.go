package metoffice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Forecast struct {
	FeelsLikeTemperature string
	Temperature string
	Weather string
	PrecipitationProbability int
}

type DayPeriod struct {

}

type WeatherLocation struct {
	Name string `json:"name"`
}

type Weather struct {
	Type string `json:"type"`
	DataDate string `json:"dataDate"`
	Location WeatherLocation `json:"Location"`
}

type Config struct {
	ApiKey string
	LocationId int
}

func GetForecast(c Config) (*Weather, error) {
	var url = fmt.Sprintf("http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/regionalforecast/json/%d?res=3hourly&key=%s", c.LocationId, c.ApiKey)

	body, err := doRequest(url)

	var raw map[string]json.RawMessage

	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw["SiteRep"], &raw)
	if err != nil {
		return nil, err
	}

	var w *Weather

	err = json.Unmarshal(raw["DV"], &w)

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