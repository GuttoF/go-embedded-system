package db

import (
	"context"
	"firebase.google.com/go"
        "firebase.google.com/go/db"
        "google.golang.org/api/option"
        "log"
	"os"
)

var FirebaseClient *db.Client

func InitFirebase() {

	databaseURL := os.Getenv("FIREBASE_DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("FIREBASE_DATABASE_URL not set")
	}

        opt := option.WithCredentialsFile("./firebase-key.json")
	config := &firebase.Config{
		DatabaseURL: databaseURL,
	}

        app, err := firebase.NewApp(context.Background(), config, opt)
        if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
        }

        client, err := app.Database(context.Background())
        if err != nil {
		log.Fatalf("error initializing firebase database: %v", err)
        }

        FirebaseClient = client
        log.Println("Connected to Firebase!")
}