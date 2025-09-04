package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetRejectionReason Возвращает причину, по которой указанному объявлению было отказано в прохождении премодерации.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetRejectionReason(ctx context.Context, accountID int, adID int) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("ad_id", strconv.Itoa(adID))

	VkRequest := types.VkRequest{
		Method: "ads.getRejectionReason",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
