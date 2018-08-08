package main

import (
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"log"
	"context"
	"github.com/davecgh/go-spew/spew"
)

func sendData(collectionName string,data []BodyData)  {
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

	for _, value := range data {
		_, _, err = client.Collection(collectionName).Add(ctx, value)
	}


	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func sendDataWithExistCheck(){
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

	var doc = client.Collection("withings_sleeps").Doc("1234")
	ref, err := doc.Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	if ref.Exists(){
		spew.Dump(ref.Data())
	} else{
		log.Println("not exist data")
	}

}