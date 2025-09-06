package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// IsNotificationsAllowed Метод проверяет, разрешил ли пользователь присылать ему уведомления в мини-приложении.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *AppMethods) IsNotificationsAllowed(ctx context.Context, userID uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("user_id", strconv.FormatUint(uint64(userID), 10))

	VkRequest := types.VkRequest{
		Method: "apps.isNotificationsAllowed",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
