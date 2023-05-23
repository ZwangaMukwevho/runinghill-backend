package repository

import "firebase.google.com/go/db"

type FirebaseRepository interface {
	AddWords() error
}

type firebaseClient struct {
	client *db.Client
}

func NewFirebaseRestClient(client *db.Client) firebaseClient {
	return firebaseClient{client}
}

func AddWords() error {
	return nil
}

// import (
// 	"fmt"
// 	"context"

// 	firebase "firebase.google.com/go"
// 	"firebase.google.com/go/auth"

// 	"google.golang.org/api/option"
//   )

//   ctx := context.Background()

//   // configure database URL
//   conf := &firebase.Config{
// 	  DatabaseURL: "https://fir-realtime-db-demo-xxxxx-default-rtdb.asia-southeast1.firebasedatabase.app",
//   }

//   // fetch service account key
//   opt := option.WithCredentialsFile("fir-realtime-db-demo-xxxxx-firebase-adminsdk-2zjoi-9b40ce1e42.json")

//   app, err := firebase.NewApp(ctx, conf, opt)
//   if err != nil {
// 	  log.Fatalln("error in initializing firebase app: ", err)
//   }

//   client, err := app.Database(ctx)
//   if err != nil {
// 	  log.Fatalln("error in creating firebase DB client: ", err)
//   }
