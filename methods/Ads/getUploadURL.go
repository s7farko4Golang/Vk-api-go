package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetUploadURLOptions struct {
	adFormat int //Формат объявления:
	//• 1 — изображение и текст.
	//• 2 — большое изображение.
	//• 3 — эксклюзивный формат.
	//• 4 — продвижение сообществ или приложений, квадратное изображение.
	//• 5 — приложение в новостной ленте (устаревший).
	//• 6 — мобильное приложение.
	//• 11 — адаптивный формат.
	//Для формата объявления «запись в сообществе» этот метод не используется,
	//т.к. фотографии являются частью записи со стены сообщества. Смотрите методы ads.createAds и wall.postAdsStealth.
	icon string //1 — получить URL для загрузки логотипа, а не основного изображения. Используется только для ad_format = 11.
}

type GetUploadURLOption func(*GetUploadURLOptions)

func GetUploadURLWithAdFormat(adFormat int) GetUploadURLOption {
	return func(o *GetUploadURLOptions) {
		o.adFormat = adFormat
	}
}

func GetUploadURLWithIcon(icon string) GetUploadURLOption {
	return func(o *GetUploadURLOptions) {
		o.icon = icon
	}
}

// GetUploadURL Возвращает URL-адрес для загрузки фотографии рекламного объявления.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetUploadURL(ctx context.Context, opts ...GetUploadURLOption) (types.VkResponse, error) {

	options := &GetUploadURLOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.adFormat != 0 {
		params.Set("ad_format", strconv.Itoa(options.adFormat))
	}
	if options.icon != "" {
		params.Set("icon", options.icon)
	}

	VkRequest := types.VkRequest{
		Method: "ads.getUploadURL",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
