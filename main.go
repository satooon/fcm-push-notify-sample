package main

import (
	"golang.org/x/net/context"

	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
	"firebase.google.com/go/messaging"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load .env: %v", err)
		os.Exit(1)
	}

	if err := os.Setenv("GCLOUD_PROJECT", os.Getenv("project_id")); err !=nil {
		log.Fatalf("error set project_id: %v", err)
		os.Exit(1)
	}

	opt := option.WithCredentialsFile(os.Getenv("credentials_file"))
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		os.Exit(1)
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error messaging client: %v", err)
		os.Exit(1)
	}

	msg := &messaging.Message{
		Topic: os.Getenv("topic"),

		//Notification: &messaging.Notification{
		//	Title: os.Getenv("title"),
		//	Body: os.Getenv("body"),
		//},

		Android: &messaging.AndroidConfig {
			Notification: &messaging.AndroidNotification{
				Title: os.Getenv("title"),
				Body: os.Getenv("body"),
			},
		},
	}

	res, err := client.Send(ctx, msg)
	if err != nil {
		log.Fatalf("error client send: %v", err)
		os.Exit(1)
	}

	log.Printf("Successfully sent message: %v\n", res)
	os.Exit(0)
}
