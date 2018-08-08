package main

import (
	"log"
	"github.com/jrmycanady/nokiahealth"
	"time"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)


type BodyData struct {
	Weight         float64
	FatFreeMass    float64
	FatRatios      float64
	FatMassWeights float64
	CreatedAt      time.Time
}

type WithingsManager struct {
	config Config
}

func (w WithingsManager)FetchWeightData(startDate,endDate time.Time) *nokiahealth.BodyMeasures{
	//var config Config
	//_, err := toml.DecodeFile("config.toml", &config)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//create client
	client := nokiahealth.NewClient(w.config.AuthData.ComsumerKey, w.config.AuthData.ComsumerSecret, "localhost")
	user := client.GenerateUser(w.config.AuthData.AccessToken, w.config.AuthData.AccessSecret, w.config.AuthData.UserID)

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

func (w WithingsManager)FetchSleepData(startDate,endDate time.Time){
	client := nokiahealth.NewClient(w.config.AuthData.ComsumerKey, w.config.AuthData.ComsumerSecret, "localhost")
	user := client.GenerateUser(w.config.AuthData.AccessToken, w.config.AuthData.AccessSecret, w.config.AuthData.UserID)

	p := nokiahealth.SleepSummaryQueryParam{}

	p.StartDateYMD = &startDate
	p.EndDateYMD = &endDate

	//get weight data
	m, err := user.GetSleepSummary(&p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Body.Series)

}

func (w WithingsManager)FetchActivityData(startDate,endDate time.Time){
	client := nokiahealth.NewClient(w.config.AuthData.ComsumerKey, w.config.AuthData.ComsumerSecret, "localhost")
	user := client.GenerateUser(w.config.AuthData.AccessToken, w.config.AuthData.AccessSecret, w.config.AuthData.UserID)

	p := nokiahealth.ActivityMeasuresQueryParam{}

	p.Date = &startDate
	p.EndDateYMD = &endDate

	//get weight data
	m, err := user.GetActivityMeasures(&p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*m.Body.Steps)
	spew.Dump(m)

}