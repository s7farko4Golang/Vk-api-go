package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type CreateTargetPixelOptions struct {
	accountId  int    //Идентификатор рекламного кабинета.
	clientId   int    //Только для рекламных агентств. ID клиента, в рекламном кабинете которого будет создаваться пиксель.
	name       string //Название пикселя — строка до 64 символов.
	domain     string //Домен сайта, на котором будет размещен пиксель.
	categoryId int    //Идентификатор категории сайта, на котором будет размещен пиксель. Для получения списка возможных идентификаторов следует использовать метод ads.getSuggestions (раздел interest_categories_v2).
}
type CreateTargetPixelOption func(*CreateTargetPixelOptions)

func CreateTargetPixelWithClientId(clientId int) CreateTargetPixelOption {
	return func(o *CreateTargetPixelOptions) {
		o.clientId = clientId
	}
}

func CreateTargetPixelWithDomain(domain string) CreateTargetPixelOption {
	return func(o *CreateTargetPixelOptions) {
		o.domain = domain
	}
}

// CreateTargetPixel Создаёт пиксель ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) CreateTargetPixel(ctx context.Context, accountID int, name string, categoryId int, opts ...CreateTargetPixelOption) (types.VkResponse, error) {

	options := &CreateTargetPixelOptions{
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
		Method: "ads.createTargetPixel",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
