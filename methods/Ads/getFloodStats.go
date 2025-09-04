package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetFloodStats Возвращает информацию о текущем состоянии счетчика — количество оставшихся запусков методов и время до
// следующего обнуления счетчика в секундах.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetFloodStats(ctx context.Context, accountID int) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))

	VkRequest := types.VkRequest{
		Method: "ads.getFloodStats",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
