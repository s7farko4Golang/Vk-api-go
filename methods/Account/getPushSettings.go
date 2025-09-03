package account_methods

import (
	"Vk-api-go/types"
	"context"
	"errors"
	"net/url"
)

type PushSettingsParams struct {
	Token string //Идентификатор устройства, используемый для отправки уведомлений.
	// (для mpns идентификатор должен представлять из себя URL для отправки уведомлений)
	DeviceID string //Уникальный идентификатор устройства.
}

// GetPushSettings Позволяет получать настройки Push-уведомлений.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) GetPushSettings(ctx context.Context, params PushSettingsParams) (types.VkResponse, error) {
	values := url.Values{}

	if params.Token != "" {
		values.Set("token", params.Token)
	} else if params.DeviceID != "" {
		values.Set("device_id", params.DeviceID)
	} else {
		return types.VkResponse{}, errors.New("either token or device_id must be provided")
	}

	vkRequest := types.VkRequest{
		Method: "account.getPushSettings",
		Params: values,
	}
	return am.methods.Call(ctx, vkRequest)
}
