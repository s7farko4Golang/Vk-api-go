package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type CreateLookalikeRequestOptions struct {
	accountId int //Идентификатор рекламного кабинета.
	clientId  int //Только для рекламных агентств.
	//Идентификатор клиента, для которого будет создаваться аудитория.
	sourceType         string //Тип источника исходной аудитории. На данный момент может принимать единственное значение retargeting_group.
	retargetingGroupId int
}
type CreateLookalikeRequestOption func(*CreateLookalikeRequestOptions)

func CreateLookalikeRequestWithClientId(clientId int) CreateLookalikeRequestOption {
	return func(o *CreateLookalikeRequestOptions) {
		o.clientId = clientId
	}
}
func CreateLookalikeRequestWithRetargetingGroupId(retargetingGroupId int) CreateLookalikeRequestOption {
	return func(o *CreateLookalikeRequestOptions) {
		o.retargetingGroupId = retargetingGroupId
	}
}

// CreateLookalikeRequest Создаёт запрос на поиск похожей аудитории.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) CreateLookalikeRequest(ctx context.Context, accountID int, sourceType string, opts ...CreateLookalikeRequestOption) (types.VkResponse, error) {

	options := &CreateLookalikeRequestOptions{
		accountId:  accountID,
		sourceType: sourceType,
	}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(options.accountId))
	params.Set("source_type", options.sourceType)
	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}

	if options.retargetingGroupId != 0 {
		params.Set("retargeting_group_id", strconv.Itoa(options.retargetingGroupId))
	}

	VkRequest := types.VkRequest{
		Method: "ads.createLookalikeRequest",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
