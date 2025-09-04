package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetTargetPixelsOptions struct {
	accountID int //Идентификатор рекламного кабинета.
	clientID  int //Только для рекламных агентств. Id клиента, в рекламном кабинете которого находятся пиксели.
}

type GetTargetPixelsOption func(*GetTargetPixelsOptions)

func GetTargetPixelsWithClientID(clientID int) GetTargetPixelsOption {
	return func(o *GetTargetPixelsOptions) {
		o.accountID = clientID
	}
}

// GetTargetPixels Возвращает список пикселей ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetTargetPixels(ctx context.Context, accountID int, opts ...GetTargetPixelsOption) (types.VkResponse, error) {

	options := &GetTargetPixelsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("accountId", strconv.Itoa(accountID))

	if options.clientID != 0 {
		params.Set("client_id", strconv.Itoa(options.clientID))
	}

	VkRequest := types.VkRequest{
		Method: "ads.getTargetPixels",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
