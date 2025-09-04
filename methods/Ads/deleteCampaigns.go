package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// DeleteCampaigns Архивирует рекламные кампании.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) DeleteCampaigns(ctx context.Context, accountID int, ids string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("ids", ids)
	VkRequest := types.VkRequest{
		Method: "ads.deleteCampaigns",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
