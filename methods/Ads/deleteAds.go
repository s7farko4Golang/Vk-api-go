package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// DeleteAds Архивирует рекламные объявления.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) DeleteAds(ctx context.Context, accountID int, ids string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("ids", ids)
	VkRequest := types.VkRequest{
		Method: "ads.deleteAds",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
