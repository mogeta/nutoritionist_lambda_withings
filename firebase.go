package main

import (
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"context"
)

func sendData(collectionName string,data interface{})  {
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


	_, _, err = client.Collection(collectionName).Add(ctx, data)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}