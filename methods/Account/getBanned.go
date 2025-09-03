package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type GetBannedOptions struct {
	fields string
	offset string //Смещение, необходимое для выборки определенного подмножества черного списка.
	count  string //Количество объектов, информацию о которых необходимо вернуть.
}
type GetBannedOption func(*GetBannedOptions)

func GetBannedWithOffset(offset string) GetBannedOption {
	return func(o *GetBannedOptions) {
		o.offset = offset
	}
}

func GetBannedWithCount(count string) GetBannedOption {
	return func(o *GetBannedOptions) {
		o.count = count
	}
}

// GetBanned Возвращает список пользователей, находящихся в черном списке.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow
func (am *AccountMethods) GetBanned(ctx context.Context, fields string, opts ...GetBannedOption) (types.VkResponse, error) {

	options := &GetBannedOptions{
		fields: fields,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("fields", options.fields)
	if options.offset != "" {
		params.Set("offset", options.offset)
	}

	if options.count != "" {
		params.Set("count", options.count)
	}

	VkRequest := types.VkRequest{
		Method: "account.getBanned",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
