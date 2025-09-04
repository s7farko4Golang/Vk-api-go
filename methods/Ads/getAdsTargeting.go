package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetAdsTargetingOptions struct {
	accountId      int  //Идентификатор рекламного кабинета.
	clientId       int  //Для рекламных агентств. Идентификатор клиента, у которого запрашиваются рекламные объявления.
	includeDeleted bool //Флаг, задающий необходимость вывода архивных объявлений.
	//• 0 — выводить только активные объявления;
	//• 1 — выводить все объявления.
	onlyDeleted bool
	campaignIds string //Фильтр по рекламным кампаниям. Сериализованный JSON-массив, содержащий ID кампаний.
	// Если параметр равен null, то будут выводиться рекламные объявления всех кампаний.
	adIds string //Фильтр по рекламным объявлениям. Сериализованный JSON-массив, содержащий ID объявлений.
	// Если параметр равен null, то будут выводиться все рекламные объявления.
	limit int //Ограничение на количество возвращаемых объявлений.
	// Используется, только если параметр ad_ids равен null, а параметр campaign_ids содержит ID только одной кампании.
	offset int //Смещение. Используется в тех же случаях, что и параметр limit.
}
type GetAdsTargetingOption func(*GetAdsTargetingOptions)

func GetAdsTargetingWithClientId(clientId int) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.clientId = clientId
	}
}

func GetAdsTargetingWithIncludeDeleted(includeDeleted bool) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.includeDeleted = includeDeleted
	}
}

func GetAdsTargetingWithOnlyDeleted(onlyDeleted bool) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.onlyDeleted = onlyDeleted
	}
}

func GetAdsTargetingWithCampaignIds(campaignIds string) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.campaignIds = campaignIds
	}
}

func GetAdsTargetingWithAdIds(adIds string) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.adIds = adIds
	}
}

func GetAdsTargetingWithLimit(limit int) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.limit = limit
	}
}

func GetAdsTargetingWithOffset(offset int) GetAdsTargetingOption {
	return func(o *GetAdsTargetingOptions) {
		o.offset = offset
	}
}

// GetAdsTargeting Возвращает параметры таргетинга рекламных объявлений
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetAdsTargeting(ctx context.Context, accountID int, opts ...GetAdsTargetingOption) (types.VkResponse, error) {

	options := &GetAdsTargetingOptions{
		accountId: accountID,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}
	if options.includeDeleted {
		params.Set("include_deleted", "1")
	}
	if options.onlyDeleted {
		params.Set("only_deleted", "1")
	}
	if options.campaignIds != "" {
		params.Set("campaign_ids", options.campaignIds)
	}
	if options.adIds != "" {
		params.Set("ad_ids", options.adIds)
	}
	if options.limit != 0 {
		params.Set("limit", strconv.Itoa(options.limit))
	}
	if options.offset != 0 {
		params.Set("offset", strconv.Itoa(options.offset))
	}

	VkRequest := types.VkRequest{
		Method: "ads.getAdsTargeting",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
