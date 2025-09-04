package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// getAccounts Возвращает список рекламных кабинетов.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) getAccounts(ctx context.Context) (types.VkResponse, error) {
	params := url.Values{}

	VkRequest := types.VkRequest{
		Method: "ads.getAccounts",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)

}
