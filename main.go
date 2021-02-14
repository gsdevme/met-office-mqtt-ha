package main

import (
	"fmt"
	"log"
	"metOfficeMqtt/metoffice"
	"os"
	"strconv"
)

func main() {
	var c metoffice.Config

	atoi, err := strconv.Atoi(os.Getenv("MET_OFFICE_LOCATION_ID"))

	if err != nil {
		log.Println("Invalid MET_OFFICE_LOCATION_ID")
		os.Exit(1)
	}

	c.ApiKey = os.Getenv("MET_OFFICE_API_KEY")
	c.LocationId = atoi

	forecast, err := metoffice.GetForecast(c)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(forecast.Location.Name)
}
