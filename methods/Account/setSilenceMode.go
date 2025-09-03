package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type SetSilenceModeOptions struct {
	token    string //Идентификатор устройства для сервиса push уведомлений.
	deviceId string //Уникальный идентификатор устройства.
	time     int    //Время в секундах на которое требуется отключить уведомления, -1 отключить навсегда.
	chatId   int    //Идентификатор чата, для которого следует отключить уведомления.
	userId   int    //Идентификатор пользователя, для которого следует отключить уведомления.
	peerId   int    //Идентификатор назначения.
	//Для пользователя:
	//id пользователя.
	//Для групповой беседы:
	//2000000000 + id беседы.
	//Для сообщества:
	//-id сообщества.
	sound int //1 — включить звук в этом диалоге, 0 — отключить звук
	// (параметр работает, только если в peer_id передан идентификатор групповой беседы или пользователя).
}
type SetSilenceModeOption func(*SetSilenceModeOptions)

func SetSilenceModeWithToken(token string) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.token = token
	}
}

func SetSilenceModeWithDeviceId(deviceId string) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.deviceId = deviceId
	}
}

func SetSilenceModeWithTime(time int) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.time = time
	}
}

func SetSilenceModeWithChatId(chatId int) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.chatId = chatId
	}
}

func SetSilenceModeWithUserId(userId int) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.userId = userId
	}
}

func SetSilenceModeWithPeerId(peerId int) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.peerId = peerId
	}
}

func SetSilenceModeWithSound(sound int) SetSilenceModeOption {
	return func(o *SetSilenceModeOptions) {
		o.sound = sound
	}
}

// SetSilenceMode Отключает push-уведомления на заданный промежуток времени.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) SetSilenceMode(ctx context.Context, opts ...SetSilenceModeOption) (types.VkResponse, error) {

	options := &SetSilenceModeOptions{}

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
	if options.time != 0 {
		params.Set("time", strconv.Itoa(options.time))
	}
	if options.chatId != 0 {
		params.Set("chat_id", strconv.Itoa(options.chatId))
	}
	if options.userId != 0 {
		params.Set("user_id", strconv.Itoa(options.userId))
	}
	if options.peerId != 0 {
		params.Set("peer_id", strconv.Itoa(options.peerId))
	}
	if options.sound != 0 {
		params.Set("sound", strconv.Itoa(options.sound))
	}

	VkRequest := types.VkRequest{
		Method: "account.setSilenceMode",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
