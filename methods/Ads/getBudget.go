package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetBudget Возвращает текущий бюджет рекламного кабинета.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetBudget(ctx context.Context, accountId int) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountId))
	VkRequest := types.VkRequest{
		Method: "ads.getBudget",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)

}
