package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetDemographics Возвращает демографическую статистику по рекламным объявлениям или кампаниям.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetDemographics(ctx context.Context, accountID int, idsType string, ids string, period string, dateFrom string, dateTo string) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("ids_type", idsType)
	params.Set("ids", ids)
	params.Set("period", period)
	params.Set("date_from", dateFrom)
	params.Set("date_to", dateTo)

	VkRequest := types.VkRequest{
		Method: "ads.getDemographics",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
