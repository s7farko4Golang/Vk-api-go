package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
)

type GetCategoriesOptions struct {
	lang string //Язык, на котором нужно вернуть результаты.
}

type GetCategoriesOption func(*GetCategoriesOptions)

func GetCategoriesWithLang(lang string) GetCategoriesOption {
	return func(o *GetCategoriesOptions) {
		o.lang = lang
	}
}

// GetCategories Позволяет получить возможные тематики рекламных объявлений.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetCategories(ctx context.Context, opts ...GetCategoriesOption) (types.VkResponse, error) {

	options := &GetCategoriesOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.lang != "" {
		params.Set("lang", options.lang)
	}

	VkRequest := types.VkRequest{
		Method: "ads.getCategories",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
