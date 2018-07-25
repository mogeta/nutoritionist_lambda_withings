package main

import (
	"github.com/jrmycanady/nokiahealth"
	"time"
	"log"
	"github.com/BurntSushi/toml"
	"fmt"
)

func main() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	//create client
	client := nokiahealth.NewClient(config.AuthData.ComsumerKey, config.AuthData.ComsumerSecret, "localhost")
	user := client.GenerateUser(config.AuthData.AccessToken, config.AuthData.AccessSecret, config.AuthData.UserID)


	t := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,loc)

	p := nokiahealth.BodyMeasuresQueryParams{}

	p.StartDate = &startDate
	p.EndDate = &endDate

	m, err := user.GetBodyMeasures(&p)
	if err != nil {
		log.Fatal(err)
	}
	measures := m.ParseData()
	fmt.Println(measures);
}
