package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetCampaignsOptions struct {
	accountId      int
	clientId       int
	includeDeleted bool
	campaignIds    string
	fields         string
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
