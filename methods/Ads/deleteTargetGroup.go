package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type DeleteTargetGroupOptions struct {
	accountId     int //Идентификатор рекламного кабинета.
	clientId      int //Id клиента, в рекламном кабинете которого будет удаляться аудитория.
	targetGroupId int //Идентификатор аудитории.
}
type DeleteTargetGroupOption func(*DeleteTargetGroupOptions)

func DeleteTargetGroupWithClientId(clientId int) DeleteTargetGroupOption {
	return func(o *DeleteTargetGroupOptions) {
		o.clientId = clientId
	}
}

// DeleteTargetGroup Удаляет аудиторию ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) DeleteTargetGroup(ctx context.Context, accountID int, targetGroupId int, opts ...DeleteTargetGroupOption) (types.VkResponse, error) {

	options := &DeleteTargetGroupOptions{
		accountId:     accountID,
		targetGroupId: targetGroupId,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))
	params.Set("target_group_id", strconv.Itoa(options.targetGroupId))

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.deleteTargetGroup",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
