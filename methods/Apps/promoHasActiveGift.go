package Apps

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type PromoHasActiveGiftOptions struct {
	promoId uint //Идентификатор промо-акции.
	userId  uint //Идентификатор пользователя. Используется только при запросе с сервисным токеном.
}

type PromoHasActiveGiftOption func(*PromoHasActiveGiftOptions)

func PromoHasActiveGiftWithUserId(userId uint) PromoHasActiveGiftOption {
	return func(o *PromoHasActiveGiftOptions) {
		o.userId = userId
	}
}

// PromoHasActiveGift Проверить есть ли у пользователя подарок в игре.
// Для вызова метода можно использовать:
// •ключ доступа пользователя
// •сервисный ключ доступа
func (am *AppMethods) PromoHasActiveGift(ctx context.Context, promoId uint, opts ...PromoHasActiveGiftOption) (types.VkResponse, error) {

	options := &PromoHasActiveGiftOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("promo_id", strconv.Itoa(int(promoId)))

	if options.userId != 0 {
		params.Set("user_id", strconv.Itoa(int(options.userId)))
	}

	VkRequest := types.VkRequest{
		Method: "apps.promoHasActiveGift",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
