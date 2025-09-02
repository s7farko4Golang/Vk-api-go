package main

import (
	"Vk-api-go/account"
	client "Vk-api-go/client"
	config "Vk-api-go/config"
	"Vk-api-go/session"
	"context"
	"fmt"
	"log"
)

func main() {

	config, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client := client.NewClient()

	session := session.NewSession(client, account.NewVkAccount(config.PrimaryAccount.AccessToken, config.PrimaryAccount.UserID))

	ctx := context.Background()

	vkResponse, err := session.Account.Ban(ctx, "649194831")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vk Response: ", vkResponse.Response)
	fmt.Println("Vk Error: ", vkResponse.Error)
}
