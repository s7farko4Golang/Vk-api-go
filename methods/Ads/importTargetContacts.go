package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type ImportTargetContactsOptions struct {
	accountID     int    //Идентификатор рекламного кабинета.
	clientID      int    //Только для рекламных агентств. Id клиента, в рекламном кабинете которого находится аудитория.
	targetGroupID int    //Идентификатор аудитории таргетинга.
	contacts      string //Список телефонов, email адресов, мобильные рекламные идентификаторы (IDFA, GAID) или
	// идентификаторов пользователей, указанных через запятую. Также принимаются их MD5-хеши или SHA256-хеши.
}

type ImportTargetContactsOption func(*ImportTargetContactsOptions)

func ImportTargetContactsWithClientID(clientID int) ImportTargetContactsOption {
	return func(o *ImportTargetContactsOptions) {
		o.clientID = clientID
	}
}

// ImportTargetContacts Импортирует список контактов рекламодателя для учета зарегистрированных во ВКонтакте пользователей в аудитории ретаргетинга.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) ImportTargetContacts(ctx context.Context, accountID int, targetGroupID int, contacts string, opts ...ImportTargetContactsOption) (types.VkResponse, error) {

	options := &ImportTargetContactsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("target_group_id", strconv.Itoa(targetGroupID))
	params.Set("contacts", contacts)

	if options.clientID != 0 {
		params.Set("client_id", strconv.Itoa(options.clientID))
	}

	VkRequest := types.VkRequest{
		Method: "ads.importTargetContacts",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
