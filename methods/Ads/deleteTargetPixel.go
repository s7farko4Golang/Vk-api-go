package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type DeleteTargetPixelOptions struct {
	accountId     int //Идентификатор рекламного кабинета.
	clientId      int //Id клиента, в рекламном кабинете которого будет удаляться аудитория.
	targetPixelId int //Идентификатор пикселя.
}
type DeleteTargetPixelOption func(*DeleteTargetPixelOptions)

func DeleteTargetPixelWithClientId(clientId int) DeleteTargetPixelOption {
	return func(o *DeleteTargetPixelOptions) {
		o.clientId = clientId
	}
}

// DeleteTargetPixel Удаляет пиксель ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) DeleteTargetPixel(ctx context.Context, accountID int, targetPixelId int, opts ...DeleteTargetPixelOption) (types.VkResponse, error) {

	options := &DeleteTargetPixelOptions{
		accountId:     accountID,
		targetPixelId: targetPixelId,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))
	params.Set("target_pixel_id", strconv.Itoa(options.targetPixelId))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.deleteTargetPixel",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
