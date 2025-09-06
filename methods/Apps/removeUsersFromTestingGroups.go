package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

// RemoveUsersFromTestingGroups Метод удаляет указанных пользователей из групп тестировщиков мини-приложения.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) RemoveUsersFromTestingGroups(ctx context.Context, userIDs string) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("user_ids", userIDs)

	VkRequest := types.VkRequest{
		Method: "apps.removeUsersFromTestingGroups",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
