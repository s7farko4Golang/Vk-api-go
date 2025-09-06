package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// GetSnippets Метод возвращает информацию о сниппетах мини-приложения или игры, созданных с помощью apps.addSnippet.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) GetSnippets(ctx context.Context) (types.VkResponse, error) {

	params := url.Values{}

	VkRequest := types.VkRequest{
		Method: "apps.getSnippets",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
