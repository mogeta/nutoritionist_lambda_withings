package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"time"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	lambda.Start(execute)
}

func execute() {

	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	withings := WithingsManager{config}

	t := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,loc)

	measureData := withings.FetchWeightData(startDate,endDate)
	body := make([]BodyData, len(measureData.Weights))
	for key := range measureData.Weights {
		body[key].Weight = measureData.Weights[key].Kgs
		body[key].FatRatios = measureData.FatRatios[key].Ratio
		body[key].FatFreeMass = measureData.FatFreeMass[key].Kgs
		body[key].FatMassWeights = measureData.FatMassWeights[key].Kgs
		body[key].CreatedAt = measureData.Weights[key].Date
	}
	sendData("body",body)
}
