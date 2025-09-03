package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type SetOnlineOptions struct {
	voip bool
}

type SetOnlineOption func(*SetOnlineOptions)

func SetOnlineWithVoip(voip bool) SetOnlineOption {
	return func(o *SetOnlineOptions) {
		o.voip = voip
	}
}

// SetOnline Помечает текущего пользователя как online на 5 минут.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) SetOnline(ctx context.Context, opts ...SetOnlineOption) (types.VkResponse, error) {

	options := &SetOnlineOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.voip {
		params.Set("voip", "1")
	}

	vkRequest := types.VkRequest{
		Method: "account.setOnline",
		Params: params,
	}
	return am.methods.Call(ctx, vkRequest)
}
