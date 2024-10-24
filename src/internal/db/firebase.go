package db

import (
    "context"
    "firebase.google.com/go"
    "firebase.google.com/go/db"
    "google.golang.org/api/option"
    "log"
)

var FirebaseClient *db.Client

func InitFirebase() {
    opt := option.WithCredentialsFile("./firebase-key.json")

    app, err := firebase.NewApp(context.Background(), nil, opt)
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