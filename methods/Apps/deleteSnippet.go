package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// DeleteSnippet Метод удаляет сниппет мини-приложения или игры.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) DeleteSnippet(ctx context.Context, id int) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	VkRequest := types.VkRequest{
		Method: "apps.deleteSnippet",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
