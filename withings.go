package main

import (
	"log"
	"github.com/BurntSushi/toml"
	"github.com/jrmycanady/nokiahealth"
	"time"
)


type BodyData struct {
	Weight         float64
	FatFreeMass    float64
	FatRatios      float64
	FatMassWeights float64
	CreatedAt      time.Time
}

func fetchWeightData(startDate,endDate time.Time) *nokiahealth.BodyMeasures{
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	//create client
	client := nokiahealth.NewClient(config.AuthData.ComsumerKey, config.AuthData.ComsumerSecret, "localhost")
	user := client.GenerateUser(config.AuthData.AccessToken, config.AuthData.AccessSecret, config.AuthData.UserID)

	p := nokiahealth.BodyMeasuresQueryParams{}
	p.StartDate = &startDate
	p.EndDate = &endDate

	//get weight data
	m, err := user.GetBodyMeasures(&p)
	if err != nil {
		log.Fatal(err)
	}

	//parse
	measures := m.ParseData()
	return measures
}