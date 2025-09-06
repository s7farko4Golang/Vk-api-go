package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// RemoveTestingGroup Метод удаляет указанную группу тестировщиков мини-приложения.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) RemoveTestingGroup(ctx context.Context, groupId int) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(groupId))

	VkRequest := types.VkRequest{
		Method: "apps.removeTestingGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
