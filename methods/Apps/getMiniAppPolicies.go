package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetMiniAppPolicies Метод получает ссылки, указанные в разделе пользовательское соглашение и политика конфиденциальности мини-приложения.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AppMethods) GetMiniAppPolicies(ctx context.Context, appID uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("app_id", strconv.Itoa(int(appID)))

	VkRequest := types.VkRequest{
		Method: "apps.getMiniAppPolicies",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
