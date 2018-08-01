package main

import (
	"log"
	"context"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(execute)
}

func execute() {
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
