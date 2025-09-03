package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
	"strings"
)

type RegisterDeviceOptions struct {
	token string //Идентификатор устройства, используемый для отправки уведомлений.
	// (для mpns идентификатор должен представлять из себя URL для отправки уведомлений).
	deviceModel   string //Строковое название модели устройства.
	deviceYear    int    //Год устройства.
	deviceId      string //Уникальный идентификатор устройства.
	systemVersion string //Строковая версия операционной системы устройства.
	noText        int    //1 — не передавать текст сообщения в push-уведомлении. 0 — (по умолчанию) текст сообщения передаётся.
	subscribe     string //Список типов уведомлений, которые следует присылать.
	settings      string //Сериализованный JSON-объект, описывающий настройки уведомлений в специальном формате.
	sandbox       int    //Флаг предназначен для iOS устройств. 1 — использовать sandbox сервер для отправки push-уведомлений, 0 — не использовать.
}
type RegisterDeviceOption func(*RegisterDeviceOptions)

func RegisterDeviceWithDeviceModel(deviceModel string) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.deviceModel = deviceModel
	}
}

func RegisterDeviceWithDeviceYear(deviceYear int) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.deviceYear = deviceYear
	}
}

func RegisterDeviceWithSystemVersion(systemVersion string) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.systemVersion = systemVersion
	}
}

func RegisterDeviceWithNoText(noText int) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.noText = noText
	}
}

func RegisterDeviceWithSubscribe(subscribes ...string) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.subscribe = strings.Join(subscribes, ",")
	}
}

func RegisterDeviceWithSettings(settings string) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.settings = settings
	}
}

func RegisterDeviceWithSandbox(sandbox int) RegisterDeviceOption {
	return func(o *RegisterDeviceOptions) {
		o.sandbox = sandbox
	}
}

// RegisterDevice Подписывает устройство на базе iOS, Android, Windows Phone или Mac на получение push-уведомлений.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AccountMethods) RegisterDevice(ctx context.Context, token string, deviceID string, opts ...RegisterDeviceOption) (types.VkResponse, error) {

	options := &RegisterDeviceOptions{
		token:    token,
		deviceId: deviceID,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("token", options.token)
	params.Set("device_id", options.deviceId)

	if options.deviceModel != "" {
		params.Set("device_model", options.deviceModel)
	}

	if options.deviceYear != 0 {
		params.Set("device_year", strconv.Itoa(options.deviceYear))
	}

	if options.systemVersion != "" {
		params.Set("system_version", options.systemVersion)
	}

	if options.noText != 0 {
		params.Set("no_text", strconv.Itoa(options.noText))
	}

	if options.sandbox != 0 {
		params.Set("sandbox", strconv.Itoa(options.sandbox))
	}

	if options.subscribe != "" {
		params.Set("subscribe", options.subscribe)
	}

	if options.settings != "" {
		params.Set("settings", options.settings)
	}

	VkRequest := types.VkRequest{
		Method: "account.registerDevice",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
