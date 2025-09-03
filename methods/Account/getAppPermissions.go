package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// GetAppPermissions Метод получает настройки пользователя вашего приложения.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AccountMethods) GetAppPermissions(ctx context.Context, userId interface{}) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("owner_id", userId.(string))
	VkRequest := types.VkRequest{
		Method: "account.getAppPermissions",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
