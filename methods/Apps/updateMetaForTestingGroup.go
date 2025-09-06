package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type UpdateMetaForTestingGroupOptions struct {
	groupId int //Необязательный параметр. Идентификатор группы тестировщиков.
	// Укажите значение параметра, чтобы создать группу тестировщиков с таким идентификатором или обновить существующую группу.
	// Если создано 25 групп тестировщиков, метод вернёт ошибку Out of limits.
	webview   string //Обязательный параметр. URL, по которому доступна версия мини-приложения для тестирования.
	name      string //Обязательный параметр. Название группы тестировщиков.
	platforms string //Обязательный параметр. Названия платформ, для которых доступна версия мини-приложения,
	// перечисленные через запятую (например, "mvk, web"). Названия платформ:
	//• mobile — мобильное приложение (Android, iOS);
	//• web — десктопная версия сайта (vk.com);
	//• mvk — мобильная версия сайта (m.vk.com).
	userIds string //Необязательный параметр. Идентификаторы пользователей, которые входят в группу тестировщиков,
	// перечисленные через запятую (например, "11, 13").
	//Если суммарное количество пользователей в группе тестировщиков превысит 1000, метод вернёт ошибку Out of limits.
}

type UpdateMetaForTestingGroupOption func(*UpdateMetaForTestingGroupOptions)

func UpdateMetaForTestingGroupWithGroupId(groupId int) UpdateMetaForTestingGroupOption {
	return func(o *UpdateMetaForTestingGroupOptions) {
		o.groupId = groupId
	}
}

func UpdateMetaForTestingGroupWithUserIds(userIds string) UpdateMetaForTestingGroupOption {
	return func(o *UpdateMetaForTestingGroupOptions) {
		o.userIds = userIds
	}
}

// UpdateMetaForTestingGroup Метод создает новую или обновляет существующую группу тестировщиков мини-приложения.
// Для вызова метода можно использовать:
// •сервисный ключ доступа
func (am *AppMethods) UpdateMetaForTestingGroup(ctx context.Context, webview string, name string, platforms string, opts ...UpdateMetaForTestingGroupOption) (types.VkResponse, error) {

	options := &UpdateMetaForTestingGroupOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("name", name)
	params.Set("platforms", platforms)
	params.Set("webview", webview)
	// Add parameters to values based on options
	if options.groupId != 0 {
		params.Set("group_id", strconv.Itoa(options.groupId))
	}
	if options.userIds != "" {
		params.Set("user_ids", options.userIds)
	}

	VkRequest := types.VkRequest{
		Method: "apps.updateMetaForTestingGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
