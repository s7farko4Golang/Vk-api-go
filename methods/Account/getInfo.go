package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strings"
)

// GetInfo Возвращает информацию о текущем аккаунте.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) GetInfo(ctx context.Context, fields ...string) (types.VkResponse, error) {

	params := url.Values{}

	if len(fields) > 0 {
		params.Set("fields", strings.Join(fields, ","))
	}

	vkRequest := types.VkRequest{
		Method: "account.getInfo",
		Params: params,
	}
	return am.methods.Call(ctx, vkRequest)
}
