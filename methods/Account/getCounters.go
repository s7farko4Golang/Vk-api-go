package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strings"
)

// GetCounters возвращает ненулевые значения счётчиков пользователя
// Для вызова метода можно использовать:
// ключ доступа пользователя
func (am *AccountMethods) GetCounters(ctx context.Context, filters ...string) (types.VkResponse, error) {
	params := url.Values{}

	if len(filters) > 0 {
		params.Set("filter", strings.Join(filters, ","))
	}

	vkRequest := types.VkRequest{
		Method: "account.getCounters",
		Params: params,
	}
	return am.methods.Call(ctx, vkRequest)
}
