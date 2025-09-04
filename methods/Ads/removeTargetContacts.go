package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type RemoveTargetContactsOptions struct {
	accountID     int    //Идентификатор рекламного кабинета.
	clientID      int    //Только для рекламных агентств. Id клиента, в рекламном кабинете которого находится аудитория.
	targetGroupID int    //Идентификатор аудитории таргетинга.
	contacts      string //Список телефонов, email адресов или идентификаторов пользователей, указанных через запятую.
	// Также принимаются их MD5-хеши.
}

type RemoveTargetContactsOption func(*RemoveTargetContactsOptions)

func RemoveTargetContactsWithClientID(clientID int) RemoveTargetContactsOption {
	return func(o *RemoveTargetContactsOptions) {
		o.clientID = clientID
	}
}

// RemoveTargetContacts Принимает запрос на исключение контактов рекламодателя из аудитории ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) RemoveTargetContacts(ctx context.Context, accountID int, targetGroupID int, contacts string, opts ...RemoveTargetContactsOption) (types.VkResponse, error) {

	options := &RemoveTargetContactsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("client_id", strconv.Itoa(targetGroupID))
	params.Set("contacts", contacts)

	if options.clientID != 0 {
		params.Set("client_id", strconv.Itoa(options.clientID))
	}

	VkRequest := types.VkRequest{
		Method: "ads.removeTargetContacts",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
