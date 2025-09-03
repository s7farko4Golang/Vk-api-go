package session

import (
	"Vk-api-go/account"
	client "Vk-api-go/client"
	"Vk-api-go/methods"
	account_methods "Vk-api-go/methods/Account"
	"Vk-api-go/methods/Ads"
)

type VkSession struct {
	client  *client.Client
	account *account.VkAccount
	methods *methods.APIMethods
	Account *account_methods.AccountMethods
	Ads     *Ads.AddMethods
}

func NewSession(client *client.Client, vkAccount *account.VkAccount) *VkSession {
	apiMethods := methods.NewAPIMethods(client, vkAccount)

	return &VkSession{
		client:  client,
		account: vkAccount,
		methods: apiMethods,
		Account: account_methods.NewAccountMethods(apiMethods),
	}
}
