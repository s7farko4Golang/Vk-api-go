package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type SetPushSettingsOptions struct {
	deviceId string //Уникальный идентификатор устройства.
	settings string //Сериализованный JSON-объект, описывающий настройки уведомлений в специальном формате.
	key      string //Ключ уведомления.
	value    string //Новое значение уведомления в специальном формате.
}
type SetPushSettingsOption func(*SetPushSettingsOptions)

func SetPushSettingsWithSettings(settings string) SetPushSettingsOption {
	return func(o *SetPushSettingsOptions) {
		o.settings = settings
	}
}

func SetPushSettingsWithKey(key string) SetPushSettingsOption {
	return func(o *SetPushSettingsOptions) {
		o.key = key
	}
}

func SetPushSettingsWithValue(value string) SetPushSettingsOption {
	return func(o *SetPushSettingsOptions) {
		o.value = value
	}
}

// SetPushSettings Изменяет настройку push-уведомлений.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) SetPushSettings(ctx context.Context, deviceID string, opts ...SetPushSettingsOption) (types.VkResponse, error) {

	options := &SetPushSettingsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.settings != "" {
		params.Set("settings", options.settings)
	}

	if options.key != "" {
		params.Set("key", options.key)
	}

	if options.value != "" {
		params.Set("value", options.value)
	}

	VkRequest := types.VkRequest{
		Method: "account.setPushSettings",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
