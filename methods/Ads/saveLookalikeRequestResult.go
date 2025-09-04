package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type SaveLookalikeRequestResultOptions struct {
	accountID int //Идентификатор рекламного кабинета.
	clientID  int //Только для рекламных агентств.
	requestID int //Идентификатор запроса на поиск похожей аудитории.
	// Получить список всех запросов на поиск похожей аудитории для данного кабинета можно с помощью ads.getLookalikeRequests.
	level int //Уровень конкретного размера похожей аудитории для сохранения.
	// Получить список всех доступных размеров аудиторий можно с помощью ads.getLookalikeRequests.
}

type SaveLookalikeRequestResultOption func(*SaveLookalikeRequestResultOptions)

func SaveLookalikeRequestResultWithClientID(clientID int) SaveLookalikeRequestResultOption {
	return func(o *SaveLookalikeRequestResultOptions) {
		o.clientID = clientID
	}
}

// SaveLookalikeRequestResult Сохраняет результат поиска похожей аудитории.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) SaveLookalikeRequestResult(ctx context.Context, accountID int, requestIDint, level int, opts ...SaveLookalikeRequestResultOption) (types.VkResponse, error) {

	options := &SaveLookalikeRequestResultOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("request_id", strconv.Itoa(options.requestID))
	params.Set("level", strconv.Itoa(level))

	if options.clientID != 0 {
		params.Set("client_id", strconv.Itoa(options.clientID))
	}

	VkRequest := types.VkRequest{
		Method: "ads.saveLookalikeRequestResult",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
