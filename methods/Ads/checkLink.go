package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type CheckLinkOptions struct {
	accountId int    //Идентификатор рекламного кабинета.
	linkType  string //Вид рекламируемого объекта:
	//community — сообщество;
	//post — запись в сообществе;
	//application — приложение ВКонтакте;
	//video — видеозапись;
	//site — внешний сайт.
	linkUrl    string //Ссылка на рекламируемый объект.
	campaignId int    //Id кампании, в которой будет создаваться объявление.
}
type CheckLinkOption func(*CheckLinkOptions)

func CheckLinkWithCampaignId(campaignId int) CheckLinkOption {
	return func(o *CheckLinkOptions) {
		o.campaignId = campaignId
	}
}

// CheckLink Проверяет ссылку на рекламируемый объект.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) CheckLink(ctx context.Context, accountID int, linkType string, linkUrl string, opts ...CheckLinkOption) (types.VkResponse, error) {

	options := &CheckLinkOptions{
		accountId: accountID,
		linkType:  linkType,
		linkUrl:   linkUrl,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))
	params.Set("link_type", options.linkType)
	params.Set("link_url", options.linkUrl)
	if options.campaignId != 0 {
		params.Set("campaign_id", strconv.Itoa(options.campaignId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.checkLink",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
