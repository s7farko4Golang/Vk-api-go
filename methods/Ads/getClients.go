package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetClients Метод возвращает список клиентов рекламного агентства. Доступен только для рекламных агентств.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetClients(ctx context.Context, accountID int) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))

	VkRequest := types.VkRequest{
		Method: "ads.getClients",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
