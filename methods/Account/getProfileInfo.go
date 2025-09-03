package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// GetProfileInfo Возвращает информацию о текущем профиле.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) GetProfileInfo(ctx context.Context) (types.VkResponse, error) {

	params := url.Values{}

	vkRequest := types.VkRequest{
		Method: "account.getProfileInfo",
		Params: params,
	}
	return am.methods.Call(ctx, vkRequest)
}
