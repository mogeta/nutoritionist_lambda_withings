package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"context"
	"github.com/BurntSushi/toml"
	"github.com/jrmycanady/nokiahealth"
	"time"
	"fmt"
)

func main() {
	lambda.Start(hello)
}

func hello(event interface{}) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("./firebase_secret_key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	log.Println("finish")
	//_, _, err = client.Collection("users").Add(ctx, iratto{
	//	Type:   event.DeviceEvent.ButtonClicked.ClickType,
	//	Created: time.Now(),
	//})
	//if err != nil {
	//	log.Fatalf("Failed adding alovelace: %v", err)
	//}

}

func fetchWeightData(){
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	//create client
	client := nokiahealth.NewClient(config.AuthData.ComsumerKey, config.AuthData.ComsumerSecret, "localhost")
	user := client.GenerateUser(config.AuthData.AccessToken, config.AuthData.AccessSecret, config.AuthData.UserID)

	//set query data
	t := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	startDate := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,loc).AddDate(0,0,-1)
	endDate := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,loc)

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
	fmt.Println(measures);
}