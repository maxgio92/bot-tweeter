package main

import (
	"fmt"
	"os"

	"github.com/maxgio92/bot-tweeter/pkg/client"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("bot-tweeter v0.1.0")
	creds := client.Credentials{
		ApiKey:       os.Getenv("API_KEY"),
		ApiSecretKey: os.Getenv("API_SECRET_KEY"),
		BearerToken:  os.Getenv("BEARER_TOKEN"),
	}

	client, err := client.GetClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
		os.Exit(1)
	}

	if client == nil {
		log.Println(fmt.Errorf("Client not valid"))
		os.Exit(1)
	}

	tweet, resp, err := client.Statuses.Update("...flying...", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
