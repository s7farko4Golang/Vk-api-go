package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type ShareTargetGroupOptions struct {
	accountID         int //Идентификатор рекламного кабинета.
	clientID          int //Только для рекламных агентств. Id клиента, в рекламном кабинете которого находится аудитория.
	targetGroupID     int //Идентификатор исходной аудитории.
	shareWithClientId int //D клиента, рекламному кабинету которого необходимо предоставить доступ к аудитории.
}

type ShareTargetGroupOption func(*ShareTargetGroupOptions)

func ShareTargetGroupWithClientID(clientID int) ShareTargetGroupOption {
	return func(o *ShareTargetGroupOptions) {
		o.clientID = clientID
	}
}

func ShareTargetGroupWithShareWithClientId(shareWithClientId int) ShareTargetGroupOption {
	return func(o *ShareTargetGroupOptions) {
		o.shareWithClientId = shareWithClientId
	}
}

// ShareTargetGroup Предоставляет доступ к аудитории ретаргетинга другому рекламному кабинету.
// В результате выполнения метода возвращается идентификатор аудитории для указанного кабинета.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) ShareTargetGroup(ctx context.Context, accountID int, targetGroupID int, opts ...ShareTargetGroupOption) (types.VkResponse, error) {

	options := &ShareTargetGroupOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("target_group_id", strconv.Itoa(targetGroupID))

	if options.clientID != 0 {
		params.Set("client_id", strconv.Itoa(options.clientID))
	}
	if options.shareWithClientId != 0 {
		params.Set("share_with_client_id", strconv.Itoa(options.shareWithClientId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.shareTargetGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
