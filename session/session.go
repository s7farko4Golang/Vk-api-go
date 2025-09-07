package session

import (
	"Vk-api-go/account"
	client "Vk-api-go/client"
	"Vk-api-go/methods"
	account_methods "Vk-api-go/methods/Account"
	"Vk-api-go/methods/Ads"
	apps "Vk-api-go/methods/Apps"
	"Vk-api-go/methods/Auth"
	"Vk-api-go/methods/Board"
)

type VkSession struct {
	client  *client.Client
	account *account.VkAccount
	methods *methods.APIMethods
	Account *account_methods.AccountMethods
	Ads     *Ads.AddMethods
	Apps    *apps.AppMethods
	Auth    *Auth.AuthMethods
	Board   *Board.BoardMethods
}

func NewSession(client *client.Client, vkAccount *account.VkAccount) *VkSession {
	apiMethods := methods.NewAPIMethods(client, vkAccount)

	return &VkSession{
		client:  client,
		account: vkAccount,
		methods: apiMethods,
		Account: account_methods.NewAccountMethods(apiMethods),
		Ads:     Ads.NewAddMethods(apiMethods),
		Apps:    apps.NewAppMethods(apiMethods),
		Auth:    Auth.NewAuthMethods(apiMethods),
		Board:   Board.NewBoardMethods(apiMethods),
	}
}
