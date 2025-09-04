package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// Возвращает информацию о музыкантах на слушателей, для которых доступно таргетирование.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetMusiciansByIds(ctx context.Context, ids int) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("ids", strconv.Itoa(ids))

	VkRequest := types.VkRequest{
		Method: "ads.getMusiciansByIds",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
