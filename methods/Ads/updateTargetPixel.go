package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type UpdateTargetPixelOptions struct {
	accountId     int    //Идентификатор рекламного кабинета.
	clientId      int    //Только для рекламных агентств. ID клиента, в рекламном кабинете которого будет создаваться пиксель.
	targetPixelId int    //Идентификатор пикселя.
	name          string //Название пикселя — строка до 64 символов.
	domain        string //Домен сайта, на котором будет размещен пиксель.
	categoryId    int    //Идентификатор категории сайта, на котором будет размещен пиксель. Для получения списка возможных идентификаторов следует использовать метод ads.getSuggestions (раздел interest_categories_v2).
}
type UpdateTargetPixelOption func(*UpdateTargetPixelOptions)

func UpdateTargetPixelWithClientId(clientId int) UpdateTargetPixelOption {
	return func(o *UpdateTargetPixelOptions) {
		o.clientId = clientId
	}
}

func UpdateTargetPixelWithDomain(domain string) UpdateTargetPixelOption {
	return func(o *UpdateTargetPixelOptions) {
		o.domain = domain
	}
}

func UpdateTargetPixelWithTargetPixelId(targetPixelId int) UpdateTargetPixelOption {
	return func(o *UpdateTargetPixelOptions) {
		o.targetPixelId = targetPixelId
	}
}

// UpdateTargetPixel Создаёт пиксель ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) UpdateTargetPixel(ctx context.Context, accountID int, name string, categoryId int, opts ...UpdateTargetPixelOption) (types.VkResponse, error) {

	options := &UpdateTargetPixelOptions{
		accountId:  accountID,
		name:       name,
		categoryId: categoryId,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))
	params.Set("name", options.name)
	params.Set("category_id", strconv.Itoa(options.categoryId))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}

	if options.domain != "" {
		params.Set("domain", options.domain)
	}

	VkRequest := types.VkRequest{
		Method: "ads.updateTargetPixel",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
