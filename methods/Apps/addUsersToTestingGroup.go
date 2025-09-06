package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// AddUsersToTestingGroup Метод добавляет указанных пользователей в группу тестировщиков мини-приложения.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) AddUsersToTestingGroup(ctx context.Context, userIDs string, groupID int) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("user_ids", userIDs)
	params.Set("group_id", strconv.Itoa(groupID))
	VkRequest := types.VkRequest{
		Method: "apps.addUsersToTestingGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
