package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// SetOffline Помечает текущего пользователя как offline (только в текущем приложении).
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) SetOffline(ctx context.Context) (types.VkResponse, error) {

	params := url.Values{}

	vkRequest := types.VkRequest{
		Method: "account.setOffline",
		Params: params,
	}
	return am.methods.Call(ctx, vkRequest)
}
