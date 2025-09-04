package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// GetVideoUploadURL Возвращает URL-адрес для загрузки видеозаписи рекламного объявления.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetVideoUploadURL(ctx context.Context) (types.VkResponse, error) {

	params := url.Values{}

	VkRequest := types.VkRequest{
		Method: "ads.getVideoUploadURL",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
