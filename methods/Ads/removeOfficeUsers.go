package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// RemoveOfficeUsers Удаляет администраторов и/или наблюдателей из рекламного кабинета.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) RemoveOfficeUsers(ctx context.Context, accountID int, ids string) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("ids", ids)

	VkRequest := types.VkRequest{
		Method: "ads.removeOfficeUsers",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
