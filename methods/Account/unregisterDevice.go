package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type UnregisterDeviceOptions struct {
	token    string //Идентификатор устройства, используемый для отправки уведомлений.
	deviceId string //Уникальный идентификатор устройства.
	sandbox  bool   //Флаг предназначен для iOS устройств. Возможные значения:
	//1 — отписать устройство, использующего sandbox сервер для отправки push-уведомлений;
	//0 — отписать устройство, не использующее sandbox сервер.
}
type UnregisterDeviceOption func(*UnregisterDeviceOptions)

func UnregisterDeviceWithToken(token string) UnregisterDeviceOption {
	return func(o *UnregisterDeviceOptions) {
		o.token = token
	}
}

func UnregisterDeviceWithDeviceId(deviceId string) UnregisterDeviceOption {
	return func(o *UnregisterDeviceOptions) {
		o.deviceId = deviceId
	}
}

func UnregisterDeviceWithSandbox(sandbox bool) UnregisterDeviceOption {
	return func(o *UnregisterDeviceOptions) {
		o.sandbox = sandbox
	}
}

// UnregisterDevice Отписывает устройство от Push уведомлений.
// Для вызова метода можно использовать:
// • ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) UnregisterDevice(ctx context.Context, opts ...UnregisterDeviceOption) (types.VkResponse, error) {

	options := &UnregisterDeviceOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.token != "" {
		params.Set("token", options.token)
	}

	if options.deviceId != "" {
		params.Set("device_id", options.deviceId)
	}

	if options.sandbox {
		params.Set("sandbox", "1")
	}

	VkRequest := types.VkRequest{
		Method: "account.unregisterDevice",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
