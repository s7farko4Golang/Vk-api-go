package Auth

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// Restore Позволяет восстановить доступ к аккаунту, используя код, полученный через SMS.
// Данный метод доступен только приложениям, имеющим доступ к Прямой авторизации.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *AuthMethods) Restore(ctx context.Context, phone string, lastName string) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("phone", phone)
	params.Set("last_name", lastName)

	VkRequest := types.VkRequest{
		Method: "auth.restore",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
