package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetLookalikeRequestsOptions struct {
	accountId   int    //Идентификатор рекламного кабинета.
	clientId    int    //Только для рекламных агентств. Идентификатор клиента, для которого возвращаются запросы.
	requestsIds string //Список идентификаторов запрашиваемых запросов через запятую. Максимальное количество идентификаторов в списке – 200.
	//Если этот параметр пуст, возвращаться будут все запросы.
	offset int //Смещение. Используется в связке с параметром limit.
	limit  int //Ограничение на количество возвращаемых запросов на поиск похожей аудитории. Используется в связке с параметром offset.
	//0 — вернуть только количество запросов в кабинете (у клиента в случае агентства).
	sortBy string //Сортировка элементов. Возможные значения:
	//• id – сортировать по возрастанию идентификаторов;
	//• update_time – сортировать по убыванию времени последнего обновления статуса.
}

type GetLookalikeRequestsOption func(*GetLookalikeRequestsOptions)

func GetLookalikeRequestsWithClientId(clientId int) GetLookalikeRequestsOption {
	return func(o *GetLookalikeRequestsOptions) {
		o.clientId = clientId
	}
}

func GetLookalikeRequestsWithRequestsIds(requestsIds string) GetLookalikeRequestsOption {
	return func(o *GetLookalikeRequestsOptions) {
		o.requestsIds = requestsIds
	}
}

func GetLookalikeRequestsWithOffset(offset int) GetLookalikeRequestsOption {
	return func(o *GetLookalikeRequestsOptions) {
		o.offset = offset
	}
}

func GetLookalikeRequestsWithLimit(limit int) GetLookalikeRequestsOption {
	return func(o *GetLookalikeRequestsOptions) {
		o.limit = limit
	}
}

func GetLookalikeRequestsWithSortBy(sortBy string) GetLookalikeRequestsOption {
	return func(o *GetLookalikeRequestsOptions) {
		o.sortBy = sortBy
	}
}

// GetLookalikeRequests Возвращает список запросов на поиск похожей аудитории.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetLookalikeRequests(ctx context.Context, opts ...GetLookalikeRequestsOption) (types.VkResponse, error) {

	options := &GetLookalikeRequestsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	if options.clientId != 0 {
		params.Set("client_id", strconv.Itoa(options.clientId))
	}
	if options.requestsIds != "" {
		params.Set("requests_ids", options.requestsIds)
	}
	if options.offset != 0 {
		params.Set("offset", strconv.Itoa(options.offset))
	}
	if options.limit != 0 {
		params.Set("limit", strconv.Itoa(options.limit))
	}
	if options.sortBy != "" {
		params.Set("sort_by", options.sortBy)
	}

	VkRequest := types.VkRequest{
		Method: "ads.getLookalikeRequests",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
