package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetCampaignsOptions struct {
	accountId      int  //Идентификатор рекламного кабинета.
	clientId       int  //Обязателен для рекламных агентств, в остальных случаях не используется. Идентификатор клиента, у которого запрашиваются рекламные кампании.
	includeDeleted bool //Флаг, задающий необходимость вывода архивных объявлений.
	//• 0 — выводить только активные кампании;
	//• 1 — выводить все кампании.
	campaignIds string //Фильтр выводимых рекламных кампаний. Сериализованный JSON-массив, содержащий id кампаний.
	// Выводиться будут только кампании, присутствующие в campaign_ids и являющиеся кампаниями указанного рекламного кабинета.
	//Если параметр равен строке null, то выводиться будут все кампании.
	fields string //Добавляет дополнительные поля в ответ. Поддерживаемые значения:
	//• ads_count — количество объявлений в кампании.
}

type GetCampaignsOption func(*GetCampaignsOptions)

func GetCampaignsWithClientId(clientId int) GetCampaignsOption {
	return func(o *GetCampaignsOptions) {
		o.clientId = clientId
	}
}

func GetCampaignsWithIncludeDeleted(includeDeleted bool) GetCampaignsOption {
	return func(o *GetCampaignsOptions) {
		o.includeDeleted = includeDeleted
	}
}

func GetCampaignsWithCampaignIds(campaignIds string) GetCampaignsOption {
	return func(o *GetCampaignsOptions) {
		o.campaignIds = campaignIds
	}
}

func GetCampaignsWithFields(fields string) GetCampaignsOption {
	return func(o *GetCampaignsOptions) {
		o.fields = fields
	}
}

// GetCampaigns Возвращает список кампаний рекламного кабинета.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetCampaigns(ctx context.Context, accountID int, opts ...GetCampaignsOption) (types.VkResponse, error) {

	options := &GetCampaignsOptions{
		accountId: accountID,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	params.Set("accountId", strconv.Itoa(options.accountId))
	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}
	if options.includeDeleted {
		params.Set("include_deleted", "1")
	}
	if options.campaignIds != "" {
		params.Set("campaign_ids", options.campaignIds)
	}
	if options.fields != "" {
		params.Set("fields", options.fields)
	}

	VkRequest := types.VkRequest{
		Method: "ads.getCampaigns",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
