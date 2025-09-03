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

	vkConfig, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	vkClient := client.NewClient()

	vkSession := session.NewSession(vkClient, account.NewVkAccount(vkConfig.PrimaryAccount.AccessToken, vkConfig.PrimaryAccount.UserID))

	ctx := context.Background()

	vkResponse, err := vkSession.Account.SetInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vk Response: ", vkResponse.Response)
	fmt.Println("Vk Error: ", vkResponse.Error)
}
