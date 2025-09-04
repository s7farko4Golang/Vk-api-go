package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetOfficeUsers Возвращает список администраторов и наблюдателей рекламного кабинета.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetOfficeUsers(ctx context.Context, accountID int) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))

	VkRequest := types.VkRequest{
		Method: "ads.getOfficeUsers",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
