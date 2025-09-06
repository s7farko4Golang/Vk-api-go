package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// DeleteAppRequests Удаляет все уведомления о запросах, отправленных из текущего приложения.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AppMethods) DeleteAppRequests(ctx context.Context) (types.VkResponse, error) {

	params := url.Values{}

	VkRequest := types.VkRequest{
		Method: "apps.deleteAppRequests",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
