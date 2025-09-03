package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type SetInfoOptions struct {
	intro          uint //Битовая маска, отвечающая за прохождение обучения в мобильных клиентах.
	ownPostDefault bool //1 – на стене пользователя по умолчанию должны отображаться только собственные записи;
	// 0 – на стене пользователя должны отображаться все записи.
	noWallReplies bool   //1 – отключить комментирование записей на стене; 0 – разрешить комментирование.
	name          string //Имя настройки
	value         string //Значение настройки
}
type SetInfoOption func(*SetInfoOptions)

func SetInfoWithIntro(intro uint) SetInfoOption {
	return func(o *SetInfoOptions) {
		o.intro = intro
	}
}

func SetInfoWithOwnPost(ownPost bool) SetInfoOption {
	return func(o *SetInfoOptions) {
		o.ownPostDefault = ownPost
	}
}

func SetInfoWithNoWallReplies(noWallReplies bool) SetInfoOption {
	return func(o *SetInfoOptions) {
		o.noWallReplies = noWallReplies
	}
}

func SetInfoWithName(name string) SetInfoOption {
	return func(o *SetInfoOptions) {
		o.name = name
	}
}

func SetInfoWithValue(value string) SetInfoOption {
	return func(o *SetInfoOptions) {
		o.value = value
	}
}

// SetInfo Позволяет редактировать информацию о текущем аккаунте.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) SetInfo(ctx context.Context, opts ...SetInfoOption) (types.VkResponse, error) {

	options := &SetInfoOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.intro != 0 {
		params.Set("device_model", strconv.FormatUint(uint64(options.intro), 10))
	}
	if options.ownPostDefault {
		params.Set("own_post_default", "1")
	}

	if options.noWallReplies {
		params.Set("no_wall_replies", "1")
	}

	if options.name != "" {
		params.Set("name", options.name)
	}

	if options.value != "" {
		params.Set("value", options.value)
	}

	VkRequest := types.VkRequest{
		Method: "account.setInfo",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
