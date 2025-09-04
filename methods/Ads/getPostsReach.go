package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// GetPostsReach Возвращает подробную статистику по охвату рекламных записей из объявлений и кампаний для продвижения записей сообщества.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetPostsReach(ctx context.Context, accountID int, idsType string, ids string) (types.VkResponse, error) {

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("ids", ids)
	params.Set("ids_type", idsType)

	VkRequest := types.VkRequest{
		Method: "ads.getPostsReach",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
