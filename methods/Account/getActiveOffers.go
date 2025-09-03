package account_methods

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type GetActiveOffersOptions struct {
	offset string //Смещение, необходимое для выборки определенного подмножества офферов.
	count  string //Количество офферов, которое необходимо получить.
}
type GetActiveOffersOption func(*GetActiveOffersOptions)

func ActiveOffersWithOffset(offset string) GetActiveOffersOption {
	return func(o *GetActiveOffersOptions) {
		o.offset = offset
	}
}

func ActiveOffersWithCount(count string) GetActiveOffersOption {
	return func(o *GetActiveOffersOptions) {
		o.count = count
	}
}

// GetActiveOffers Возвращает список активных рекламных предложений (офферов),
// выполнив которые, пользователь сможет получить соответствующее количество голосов на свой счёт внутри приложения.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
func (am *AccountMethods) GetActiveOffers(ctx context.Context, opts ...GetActiveOffersOption) (types.VkResponse, error) {

	options := &GetActiveOffersOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.offset != "" {
		params.Set("offset", options.offset)
	}

	if options.count != "" {
		params.Set("count", options.count)
	}

	VkRequest := types.VkRequest{
		Method: "account.getActiveOffers",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
