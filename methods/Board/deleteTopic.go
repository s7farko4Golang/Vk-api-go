package Board

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

// DeleteTopic Метод удаляет обсуждение в сообществе.
// Для вызова метода можно использовать:
// •ключ доступа пользователя, полученный в Standalone‑приложении через Implicit Flow (требуются права доступа: groups)
func (am *BoardMethods) DeleteTopic(ctx context.Context, groupId uint, topicId uint) (types.VkResponse, error) {

	params := url.Values{}
	params.Set("group_id", strconv.Itoa(int(groupId)))
	params.Set("topic_id", strconv.Itoa(int(topicId)))

	VkRequest := types.VkRequest{
		Method: "board.deleteTopic",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}
