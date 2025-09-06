package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type GetScopesOptions struct {
	getScopesType string //Необязательный параметр. Тип области видимости (scope). Возможные значения:
	// • user — пользователь.
	//• group — сообщество.
}

type GetScopesOption func(*GetScopesOptions)

func GetScopesWithGetScopesType(getScopesType string) GetScopesOption {
	return func(o *GetScopesOptions) {
		o.getScopesType = getScopesType
	}
}

// GetScopes Метод получает права доступа.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AppMethods) GetScopes(ctx context.Context, opts ...GetScopesOption) (types.VkResponse, error) {

	options := &GetScopesOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.getScopesType != "" {
		params.Set("type", options.getScopesType)
	}

	VkRequest := types.VkRequest{
		Method: "apps.getScopes",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
