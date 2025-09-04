package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// GetMusicians Возвращает информацию о музыкантах, на слушателей которых доступно таргетирование.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetMusicians(ctx context.Context, artistName string) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("artist_name", artistName)

	VkRequest := types.VkRequest{
		Method: "ads.getMusicians",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
