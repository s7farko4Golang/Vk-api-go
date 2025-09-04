package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetTargetingStatsOptions struct {
	accountId int
	clientId  int
}

type GetTargetingStatsOption func(*GetTargetingStatsOptions)

func GetTargetingStatsWithClientId(clientId int) GetTargetingStatsOption {
	return func(o *GetTargetingStatsOptions) {
		o.clientId = clientId
	}
}

// GetTargetingStats Возвращает список пикселей ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetTargetingStats(ctx context.Context, accountID int, opts ...GetTargetingStatsOption) (types.VkResponse, error) {

	options := &GetTargetingStatsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.getTargetingStats",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
