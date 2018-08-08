package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
	"fmt"
	"time"
	"github.com/BurntSushi/toml"
	"log"
)

func TestExampleSuccess(t *testing.T) {

	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	withings := WithingsManager{config}

	timeNow := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc)

	measureData := withings.FetchWeightData(startDate,endDate)
	body := make([]BodyData, len(measureData.Weights))
	for key := range measureData.Weights {
		body[key].Weight = measureData.Weights[key].Kgs
		body[key].FatRatios = measureData.FatRatios[key].Ratio
		body[key].FatFreeMass = measureData.FatFreeMass[key].Kgs
		body[key].FatMassWeights = measureData.FatMassWeights[key].Kgs
		body[key].CreatedAt = measureData.Weights[key].Date
	}
	fmt.Println(body)
}

func TestFetchSleepData(t *testing.T){
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	withings := WithingsManager{config}

	timeNow := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc)

	withings.FetchSleepData(startDate,endDate)
}

func TestFetchActivityData(t *testing.T){
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	withings := WithingsManager{config}

	timeNow := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc)

	withings.FetchActivityData(startDate,endDate)
}