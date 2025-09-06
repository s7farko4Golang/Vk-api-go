package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetTestingGroupsOptions struct {
	groupId int //Необязательный параметр. Идентификатор группы тестировщиков.
	// Чтобы получить идентификаторы всех групп тестировщиков мини-приложения, оставьте параметр groupId пустым.
	// Чтобы получить информацию о конкретной группе тестировщиков, укажите её идентификатор в параметре groupId.
}

type GetTestingGroupsOption func(*GetTestingGroupsOptions)

func GetTestingGroupsWithGroupId(groupId int) GetTestingGroupsOption {
	return func(o *GetTestingGroupsOptions) {
		o.groupId = groupId
	}
}

// GetTestingGroups Метод возвращает группы тестировщиков мини-приложения.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) GetTestingGroups(ctx context.Context, opts ...GetTestingGroupsOption) (types.VkResponse, error) {

	options := &GetTestingGroupsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.groupId != 0 {
		params.Set("group_id", strconv.Itoa(options.groupId))
	}

	VkRequest := types.VkRequest{
		Method: "apps.getTestingGroups",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
