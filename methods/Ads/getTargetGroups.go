package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetTargetGroupsOptions struct {
	accountId int  //Идентификатор рекламного кабинета.
	clientId  int  //Только для рекламных агентств. Id клиента, в рекламном кабинете которого находятся аудитории.
	extended  bool //Если 1, в результатах будет указан код для размещения на сайте.
	//Устаревший параметр. Используется только для старых групп ретаргетинга,
	//которые пополнялись без помощи пикселя. Для новых аудиторий его следует опускать.
}

type GetTargetGroupsOption func(*GetTargetGroupsOptions)

func GetTargetGroupsWithClientId(clientId int) GetTargetGroupsOption {
	return func(o *GetTargetGroupsOptions) {
		o.clientId = clientId
	}
}
func GetTargetGroupsWithExtended(extended bool) GetTargetGroupsOption {
	return func(o *GetTargetGroupsOptions) {
		o.extended = extended
	}
}

// GetTargetGroups Возвращает список аудиторий ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetTargetGroups(ctx context.Context, accountId int, opts ...GetTargetGroupsOption) (types.VkResponse, error) {

	options := &GetTargetGroupsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountId))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}
	if options.extended {
		params.Set("extended", "1")
	}

	VkRequest := types.VkRequest{
		Method: "ads.getTargetGroups",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
