package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetAdsOptions struct {
	accountId      int
	clientId       int
	includeDeleted bool
	onlyDeleted    bool
	campaignIds    string
	adIds          string
	limit          int
	offset         int
}
type GetAdsOption func(*GetAdsOptions)

func GetAdsWithClientId(clientId int) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.clientId = clientId
	}
}

func GetAdsWithIncludeDeleted(includeDeleted bool) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.includeDeleted = includeDeleted
	}
}

func GetAdsWithOnlyDeleted(onlyDeleted bool) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.onlyDeleted = onlyDeleted
	}
}

func GetAdsWithCampaignIds(campaignIds string) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.campaignIds = campaignIds
	}
}

func GetAdsWithAdIds(adIds string) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.adIds = adIds
	}
}

func GetAdsWithLimit(limit int) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.limit = limit
	}
}

func GetAdsWithOffset(offset int) GetAdsOption {
	return func(o *GetAdsOptions) {
		o.offset = offset
	}
}

// GetAds Возвращает список рекламных объявлений.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetAds(ctx context.Context, accountID int, linkType string, linkUrl string, opts ...GetAdsOption) (types.VkResponse, error) {

	options := &GetAdsOptions{}

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
		Method: "ads.getAds",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
