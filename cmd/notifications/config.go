// package main

// import (
// 	"context"
// 	"log"
// 	"path/filepath"

// 	firebase "firebase.google.com/go"
// 	"firebase.google.com/go/messaging"
// 	"google.golang.org/api/option"
// )

// func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

// 	ctx := context.Background()

// 	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
// 	if err != nil {
// 		log.Println("Unable to load serviceAccountKeys.json file")
// 	}

// 	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

// 	//Firebase admin SDK initialization
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Println("Firebase load error")
// 	}

// 	//Messaging client
// 	client, _ := app.Messaging(ctx)

// 	return app, ctx, client
// }
