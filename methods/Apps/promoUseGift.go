package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type PromoUseGiftOptions struct {
	promoId uint //Идентификатор промо-акции.
	userId  uint //Идентификатор пользователя. Используется только при запросе с сервисным ключом доступа.
}

type PromoUseGiftOption func(*PromoUseGiftOptions)

func PromoUseGiftWithUserId(userId uint) PromoUseGiftOption {
	return func(o *PromoUseGiftOptions) {
		o.userId = userId
	}
}

// PromoUseGift Метод отмечает подарок, полученный пользователем в промоакции, как использованный.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *AppMethods) PromoUseGift(ctx context.Context, promoId uint, opts ...PromoUseGiftOption) (types.VkResponse, error) {

	options := &PromoUseGiftOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("promo_id", strconv.Itoa(int(promoId)))

	if options.userId != 0 {
		params.Set("user_id", strconv.Itoa(int(options.userId)))
	}

	VkRequest := types.VkRequest{
		Method: "apps.promoUseGift",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
