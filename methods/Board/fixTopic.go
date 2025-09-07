package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// FixTopic Метод закрепляет обсуждение в сообществе. Закреплённое обсуждение при любой сортировке выводится первым в списке.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) FixTopic(ctx context.Context, groupId, topicId uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))

	VkRequest := types.VkRequest{
		Method: "board.fixTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
