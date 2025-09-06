package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type SendRequestOptions struct {
	userId          uint   //Идентификатор пользователя, которому следует отправить запрос.
	text            string //Текст запроса.
	sendRequestType string //Тип запроса. Возможные значения:
	//• invite – если запрос отправляется пользователю, не установившему приложение;
	//• request – если пользователь уже установил приложение.
	//Обратите внимание! Для запросов с type = invite действует ограничение — одному и тому же
	//пользователю нельзя отправить запрос чаще одного раза в неделю.
	name string //Уникальное в рамках приложения имя для каждого вида отправляемого запроса.
	// Макс. длина = 128
	key string //Строка, которая будет возвращена назад при переходе пользователя по запросу в приложение.
	// Может использоваться для подсчета конверсии.
	separate bool //Запрет на группировку запроса с другими, имеющими тот же name.
	// 1 — группировка запрещена,
	// 0 — группировка разрешена.
	// По умолчанию: 0.
}

type SendRequestOption func(*SendRequestOptions)

func SendRequestWithText(text string) SendRequestOption {
	return func(o *SendRequestOptions) {
		o.text = text
	}
}
func SendRequestWithType(sendRequestType string) SendRequestOption {
	return func(o *SendRequestOptions) {
		o.sendRequestType = sendRequestType
	}
}
func SendRequestWithName(name string) SendRequestOption {
	return func(o *SendRequestOptions) {
		o.name = name
	}
}
func SendRequestWithKey(key string) SendRequestOption {
	return func(o *SendRequestOptions) {
		o.key = key
	}
}
func SendRequestWithSeparate(separate bool) SendRequestOption {
	return func(o *SendRequestOptions) {
		o.separate = separate
	}
}

// SendRequest Позволяет отправить запрос другому пользователю в приложении, использующем авторизацию ВКонтакте.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AppMethods) SendRequest(ctx context.Context, userId uint, opts ...SendRequestOption) (types.VkResponse, error) {

	options := &SendRequestOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("user_id", strconv.FormatUint(uint64(userId), 10))

	if options.text != "" {
		params.Set("text", options.text)
	}
	if options.sendRequestType != "" {
		params.Set("type", options.sendRequestType)
	}
	if options.name != "" {
		params.Set("name", options.name)
	}
	if options.key != "" {
		params.Set("key", options.key)
	}
	if options.separate {
		params.Set("separate", "1")
	}

	VkRequest := types.VkRequest{
		Method: "apps.sendRequest",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
