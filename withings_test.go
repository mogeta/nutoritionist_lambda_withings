package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
	"fmt"
	"time"
)

func TestExampleSuccess(t *testing.T) {
	
	timeNow := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(timeNow.Year(),timeNow.Month(),timeNow.Day(),0,0,0,0,loc)

	measureData := fetchWeightData(startDate,endDate)
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
